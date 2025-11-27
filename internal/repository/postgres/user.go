package postgres

import (
	"errors"

	"github.com/antongoncharik/sso/internal/domain"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type UserRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) domain.UserRepo {
	return &UserRepo{db: db}
}

func (r *UserRepo) GetUserByID(id int64) (domain.User, error) {
	userData := []domain.User{}

	err := r.db.Select(&userData, "SELECT id, email FROM users WHERE id = $1", id)

	if len(userData) == 0 {
		return domain.User{}, errors.New("user not found")
	}

	return userData[0], err
}

func (r *UserRepo) GetUserByEmail(email string) (domain.User, error) {
	userData := []domain.User{}

	err := r.db.Select(&userData, "SELECT id, password_hash FROM users WHERE email = $1", email)

	if len(userData) == 0 {
		return domain.User{}, errors.New("user not found")
	}

	return userData[0], err
}

func (r *UserRepo) CreateUser(user domain.User) error {
	_, err := r.db.Exec(`INSERT INTO users (email, password_hash)
	VALUES ($1, $2)`, user.Email, user.PasswordHash)

	return err
}
