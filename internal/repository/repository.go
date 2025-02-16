package repository

import (
	"errors"

	"github.com/antongoncharik/sso/internal/entity"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Repository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) GetUserByID(id int64) (entity.User, error) {
	userData := []entity.User{}

	err := r.db.Select(&userData, "SELECT id, email FROM users WHERE id = $1", id)

	if len(userData) == 0 {
		return entity.User{}, errors.New("user not found")
	}

	return userData[0], err
}

func (r *Repository) GetUserByEmail(email string) (entity.User, error) {
	userData := []entity.User{}

	err := r.db.Select(&userData, "SELECT id, password_hash FROM users WHERE email = $1", email)

	if len(userData) == 0 {
		return entity.User{}, errors.New("user not found")
	}

	return userData[0], err
}

func (r *Repository) CreateUser(user entity.User) error {
	_, err := r.db.Exec(`INSERT INTO users (email, password_hash)
	VALUES ($1, $2)`, user.Email, user.PasswordHash)

	return err
}

func (r *Repository) GetClient(app string) (entity.Client, error) {
	clienData := []entity.Client{}

	err := r.db.Select(&clienData, "SELECT id, redirect_uri FROM clients WHERE app = $1", app)

	if len(clienData) == 0 {
		return entity.Client{}, errors.New("client not found")
	}

	return clienData[0], err
}

func (r *Repository) CreateCode(code entity.Code) error {
	_, err := r.db.Exec(`INSERT INTO codes (code, exp, user_id, client_id)
	VALUES ($1, $2, $3, $4)`, code.Code, code.Exp, code.UserId, code.ClientId)

	return err
}

func (r *Repository) CreateToken(createToken entity.CreateToken) error {
	_, err := r.db.Exec(`INSERT INTO tokens (user_id, access_token, refresh_token)
	VALUES ($1, $2, $3)`, createToken.UserId, createToken.AccessToken, createToken.RefreshToken)

	return err
}

func (r *Repository) GetCode(exchangeCode entity.ExchangeCode) (entity.Code, error) {
	codeData := []entity.Code{}

	err := r.db.Select(&codeData, "SELECT id, user_id, exp FROM codes WHERE code = $1", exchangeCode.Code)

	if len(codeData) == 0 {
		return entity.Code{}, errors.New("code not found")
	}

	return codeData[0], err
}

func (r *Repository) GetToken(token string) (entity.Token, error) {
	tokenData := []entity.Token{}

	err := r.db.Select(&tokenData, "SELECT id, user_id, access_token, refresh_token FROM tokens WHERE refresh_token = $1", token)

	if len(tokenData) == 0 {
		return entity.Token{}, errors.New("token not found")
	}

	return tokenData[0], err
}
