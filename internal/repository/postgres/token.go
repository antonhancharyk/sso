package postgres

import (
	"errors"

	"github.com/antongoncharik/sso/internal/domain"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type TokenRepo struct {
	db *sqlx.DB
}

func NewTokenRepo(db *sqlx.DB) domain.TokenRepo {
	return &TokenRepo{db: db}
}

func (r *TokenRepo) CreateToken(createToken domain.CreateToken) error {
	_, err := r.db.Exec(`INSERT INTO tokens (user_id, access_token, refresh_token)
	VALUES ($1, $2, $3)`, createToken.UserId, createToken.AccessToken, createToken.RefreshToken)

	return err
}

func (r *TokenRepo) GetToken(token string) (domain.Token, error) {
	tokenData := []domain.Token{}

	err := r.db.Select(&tokenData, "SELECT id, user_id, access_token, refresh_token FROM tokens WHERE refresh_token = $1", token)

	if len(tokenData) == 0 {
		return domain.Token{}, errors.New("token not found")
	}

	return tokenData[0], err
}
