package service

import (
	"errors"
	"time"

	"github.com/antongoncharik/sso/internal/entity"
	"github.com/antongoncharik/sso/internal/repository"
	"github.com/antongoncharik/sso/pkg/utilities"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	repo *repository.Repository
	keys entity.RSAKeys
}

func New(repo *repository.Repository, keys entity.RSAKeys) *Service {
	return &Service{repo, keys}
}

func (s *Service) Register(register entity.Register) (entity.Code, error) {
	emailAllowed := false
	emails := []string{"ant.goncharik@gmail.com", "bnncrmknt@gmail.com"}
	for _, v := range emails {
		if register.Email == v {
			emailAllowed = true
		}
	}
	if !emailAllowed {
		return entity.Code{}, errors.New("email is not allowed")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(register.Password), bcrypt.DefaultCost)
	if err != nil {
		return entity.Code{}, err
	}

	err = s.repo.CreateUser(entity.User{Email: register.Email, PasswordHash: string(hash)})
	if err != nil {
		return entity.Code{}, err
	}

	user, err := s.repo.GetUserByEmail(register.Email)
	if err != nil {
		return entity.Code{}, err
	}

	client, err := s.repo.GetClient(register.App)
	if err != nil {
		return entity.Code{}, err
	}

	if client.RedirectURI != register.RedirectUri {
		return entity.Code{}, errors.New("invalId redirect URI")
	}

	exp := time.Now().Add(5 * time.Minute)

	code := uuid.New()

	err = s.repo.CreateCode(entity.Code{Code: code.String(), Exp: exp, UserId: user.Id, ClientId: client.Id})
	if err != nil {
		return entity.Code{}, err
	}

	return entity.Code{Code: code.String(), UserId: user.Id, ClientId: client.Id}, nil
}

func (s *Service) Login(login entity.Login) (entity.Code, error) {
	user, err := s.repo.GetUserByEmail(login.Email)
	if err != nil {
		return entity.Code{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(login.Password))
	if err != nil {
		return entity.Code{}, errors.New("invalid password")
	}

	client, err := s.repo.GetClient(login.App)
	if err != nil {
		return entity.Code{}, err
	}

	if client.RedirectURI != login.RedirectUri {
		return entity.Code{}, errors.New("invalid redirect uri")
	}

	exp := time.Now().Add(5 * time.Minute)

	code := uuid.New()

	err = s.repo.CreateCode(entity.Code{Code: code.String(), Exp: exp, UserId: user.Id, ClientId: client.Id})
	if err != nil {
		return entity.Code{}, err
	}

	return entity.Code{Code: code.String(), UserId: user.Id, ClientId: client.Id}, nil
}

func (s *Service) ExchangeCode(exchangeCode entity.ExchangeCode) (entity.Token, error) {
	code, err := s.repo.GetCode(exchangeCode)
	if err != nil {
		return entity.Token{}, err
	}

	if code.Exp.Before(time.Now()) {
		return entity.Token{}, errors.New("code has expired")
	}

	user, err := s.repo.GetUserByID(code.UserId)
	if err != nil {
		return entity.Token{}, err
	}

	token, err := utilities.GenerateToken(code.UserId, user.Email, s.keys.PrivateKey)
	if err != nil {
		return entity.Token{}, err
	}

	err = s.repo.CreateToken(entity.CreateToken{UserId: code.UserId, AccessToken: token.AccessToken, RefreshToken: token.RefreshToken})
	if err != nil {
		return entity.Token{}, err
	}

	return token, nil
}

func (s *Service) RefreshToken(validateToken entity.ValidateToken) (entity.Token, error) {
	token, err := s.repo.GetToken(validateToken.Token)
	if err != nil {
		return entity.Token{}, err
	}

	user, err := s.repo.GetUserByID(token.UserId)
	if err != nil {
		return entity.Token{}, err
	}

	token, err = utilities.GenerateToken(token.UserId, user.Email, s.keys.PrivateKey)
	if err != nil {
		return entity.Token{}, err
	}

	err = s.repo.CreateToken(entity.CreateToken{UserId: token.UserId, AccessToken: token.AccessToken, RefreshToken: token.RefreshToken})
	if err != nil {
		return entity.Token{}, err
	}

	return token, nil
}

func (s *Service) ValidateToken(token string) error {
	return utilities.ValidateToken(token, s.keys.PublicKey)
}
