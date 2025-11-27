package service

import (
	"errors"
	"time"

	"github.com/antongoncharik/sso/internal/domain"
	"github.com/antongoncharik/sso/internal/infra/security"
)

type CodeService struct {
	userRepo  domain.UserRepo
	tokenRepo domain.TokenRepo
	codeRepo  domain.CodeRepo
	RSA       security.RSA
}

func NewCodeService(userRepo domain.UserRepo, tokenRepo domain.TokenRepo, codeRepo domain.CodeRepo, RSA security.RSA) *CodeService {
	return &CodeService{userRepo: userRepo, tokenRepo: tokenRepo, codeRepo: codeRepo}
}

func (s *CodeService) ExchangeCode(exchangeCode domain.ExchangeCode) (domain.Token, error) {
	code, err := s.codeRepo.GetCode(exchangeCode)
	if err != nil {
		return domain.Token{}, err
	}

	if code.Exp.Before(time.Now()) {
		return domain.Token{}, errors.New("code has expired")
	}

	user, err := s.userRepo.GetUserByID(code.UserId)
	if err != nil {
		return domain.Token{}, err
	}

	token, err := security.GenerateToken(code.UserId, user.Email, s.RSA.PrivateKey)
	if err != nil {
		return domain.Token{}, err
	}

	err = s.tokenRepo.CreateToken(domain.CreateToken{UserId: code.UserId, AccessToken: token.AccessToken, RefreshToken: token.RefreshToken})
	if err != nil {
		return domain.Token{}, err
	}

	return token, nil
}
