package repository

import (
	"github.com/antongoncharik/sso/internal/domain"
	"github.com/antongoncharik/sso/internal/repository/postgres"
	"github.com/jmoiron/sqlx"
)

type Repositories struct {
	User   domain.UserRepo
	Client domain.ClientRepo
	Token  domain.TokenRepo
	Code   domain.CodeRepo
}

func NewRepositories(db *sqlx.DB) *Repositories {
	return &Repositories{
		User:   postgres.NewUserRepo(db),
		Client: postgres.NewClientRepo(db),
		Token:  postgres.NewTokenRepo(db),
		Code:   postgres.NewCodeRepo(db),
	}
}
