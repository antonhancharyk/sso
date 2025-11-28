package service

import (
	"errors"
	"time"

	"github.com/antongoncharik/sso/internal/domain"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepo   domain.UserRepo
	codeRepo   domain.CodeRepo
	clientRepo domain.ClientRepo
}

func NewUserService(userRepo domain.UserRepo, codeRepo domain.CodeRepo, clientRepo domain.ClientRepo) *UserService {
	return &UserService{userRepo: userRepo, codeRepo: codeRepo, clientRepo: clientRepo}
}

func (s *UserService) Register(register domain.Register) (domain.Code, error) {
	emailAllowed := false
	emails := []string{"ant.goncharik@gmail.com", "bnncrmknt@gmail.com"}
	for _, v := range emails {
		if register.Email == v {
			emailAllowed = true
		}
	}
	if !emailAllowed {
		return domain.Code{}, errors.New("email is not allowed")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(register.Password), bcrypt.DefaultCost)
	if err != nil {
		return domain.Code{}, err
	}

	err = s.userRepo.CreateUser(domain.User{Email: register.Email, PasswordHash: string(hash)})
	if err != nil {
		return domain.Code{}, err
	}

	user, err := s.userRepo.GetUserByEmail(register.Email)
	if err != nil {
		return domain.Code{}, err
	}

	client, err := s.clientRepo.GetClient(register.App)
	if err != nil {
		return domain.Code{}, err
	}

	if client.RedirectURI != register.RedirectUri {
		return domain.Code{}, errors.New("invalId redirect URI")
	}

	exp := time.Now().Add(5 * time.Minute)

	code := uuid.New()

	err = s.codeRepo.CreateCode(domain.Code{Code: code.String(), Exp: exp, UserId: user.Id, ClientId: client.Id})
	if err != nil {
		return domain.Code{}, err
	}

	return domain.Code{Code: code.String(), UserId: user.Id, ClientId: client.Id}, nil
}

func (s *UserService) Login(login domain.Login) (domain.Code, error) {
	user, err := s.userRepo.GetUserByEmail(login.Email)
	if err != nil {
		return domain.Code{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(login.Password))
	if err != nil {
		return domain.Code{}, errors.New("invalid password")
	}

	client, err := s.clientRepo.GetClient(login.App)
	if err != nil {
		return domain.Code{}, err
	}

	if client.RedirectURI != login.RedirectUri {
		return domain.Code{}, errors.New("invalid redirect uri")
	}

	exp := time.Now().Add(5 * time.Minute)

	code := uuid.New()

	err = s.codeRepo.CreateCode(domain.Code{Code: code.String(), Exp: exp, UserId: user.Id, ClientId: client.Id})
	if err != nil {
		return domain.Code{}, err
	}

	return domain.Code{Code: code.String(), UserId: user.Id, ClientId: client.Id}, nil
}
