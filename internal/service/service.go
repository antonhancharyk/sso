package service

import (
	"github.com/antongoncharik/sso/internal/domain"
	"github.com/antongoncharik/sso/internal/infra/security"
)

type Service struct {
	User  *UserService
	Token *TokenService
	Code  *CodeService
}

type ServiceDeps struct {
	UserRepo   domain.UserRepo
	ClientRepo domain.ClientRepo
	CodeRepo   domain.CodeRepo
	TokenRepo  domain.TokenRepo
	RSA        security.RSA
}

func NewServices(deps ServiceDeps) *Service {
	return &Service{
		User:  NewUserService(deps.UserRepo, deps.CodeRepo),
		Token: NewTokenService(deps.TokenRepo, deps.UserRepo, deps.RSA),
		Code:  NewCodeService(deps.UserRepo, deps.TokenRepo, deps.CodeRepo, deps.RSA),
	}
}
