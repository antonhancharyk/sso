package security

import (
	"crypto/rsa"
	"errors"
	"fmt"
	"time"

	"github.com/antongoncharik/sso/internal/domain"
	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(userId int64, email string, privateKey *rsa.PrivateKey) (domain.Token, error) {
	accessToken := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"user_id": userId,
		"email":   email,
		"exp":     time.Now().Add(15 * time.Minute).Unix(),
	})
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"user_id": userId,
		"email":   email,
		"exp":     time.Now().Add(7 * 24 * time.Hour).Unix(),
	})

	accessTokenStr, err := accessToken.SignedString(privateKey)
	if err != nil {
		return domain.Token{}, errors.New("error generating token")
	}
	refreshTokenStr, err := refreshToken.SignedString(privateKey)
	if err != nil {
		return domain.Token{}, errors.New("error generating token")
	}

	return domain.Token{UserId: userId, AccessToken: accessTokenStr, RefreshToken: refreshTokenStr}, nil
}

func ValidateToken(tokenStr string, publicKey *rsa.PublicKey) error {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodRSA)
		if !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return publicKey, nil
	})
	if err != nil {
		return err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		exp := int64(claims["exp"].(float64))
		if exp < time.Now().Unix() {
			return errors.New("token has expired")
		}
		return nil
	}

	return errors.New("invalId token")
}
