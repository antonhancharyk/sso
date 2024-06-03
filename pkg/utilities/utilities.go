package utilities

import (
	"crypto/rsa"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/antongoncharik/sso/internal/entity"
	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(userId int64, privateKey *rsa.PrivateKey) (entity.Token, error) {
	accessToken := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"user_id": userId,
		"exp":     time.Now().Add(15 * time.Minute).Unix(),
	})
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"user_id": userId,
		"exp":     time.Now().Add(7 * 24 * time.Hour).Unix(),
	})

	accessTokenStr, err := accessToken.SignedString(privateKey)
	if err != nil {
		return entity.Token{}, errors.New("error generating token")
	}
	refreshTokenStr, err := refreshToken.SignedString(privateKey)
	if err != nil {
		return entity.Token{}, errors.New("error generating token")
	}

	return entity.Token{AccessToken: accessTokenStr, RefreshToken: refreshTokenStr}, nil
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

func LoadPrivateKey(path string) (*rsa.PrivateKey, error) {
	privateKeyData, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateKeyData)
	if err != nil {
		return nil, err
	}

	return privateKey, nil
}

func LoadPublicKey(path string) (*rsa.PublicKey, error) {
	publicKeyData, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(publicKeyData)
	if err != nil {
		return nil, err
	}

	return publicKey, nil
}
