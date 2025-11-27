package service

import (
	"github.com/antongoncharik/sso/internal/domain"
	"github.com/antongoncharik/sso/internal/infra/security"
)

type TokenService struct {
	tokenRepo domain.TokenRepo
	userRepo  domain.UserRepo
	RSA       security.RSA
}

func NewTokenService(tokenRepo domain.TokenRepo, userRepo domain.UserRepo, RSA security.RSA) *TokenService {
	return &TokenService{tokenRepo: tokenRepo, userRepo: userRepo}
}

func (s *TokenService) RefreshToken(validateToken domain.ValidateToken) (domain.Token, error) {
	token, err := s.tokenRepo.GetToken(validateToken.Token)
	if err != nil {
		return domain.Token{}, err
	}

	user, err := s.userRepo.GetUserByID(token.UserId)
	if err != nil {
		return domain.Token{}, err
	}

	token, err = security.GenerateToken(token.UserId, user.Email, s.RSA.PrivateKey)
	if err != nil {
		return domain.Token{}, err
	}

	err = s.tokenRepo.CreateToken(domain.CreateToken{UserId: token.UserId, AccessToken: token.AccessToken, RefreshToken: token.RefreshToken})
	if err != nil {
		return domain.Token{}, err
	}

	return token, nil
}

func (s *TokenService) ValidateToken(token string) error {
	return security.ValidateToken(token, s.RSA.PublicKey)
}
