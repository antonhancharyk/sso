package postgres

import (
	"errors"

	"github.com/antongoncharik/sso/internal/domain"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type ClientRepo struct {
	db *sqlx.DB
}

func NewClientRepo(db *sqlx.DB) domain.ClientRepo {
	return &ClientRepo{db: db}
}

func (r *ClientRepo) GetClient(app string) (domain.Client, error) {
	clienData := []domain.Client{}

	err := r.db.Select(&clienData, "SELECT id, redirect_uri FROM clients WHERE app = $1", app)

	if len(clienData) == 0 {
		return domain.Client{}, errors.New("client not found")
	}

	return clienData[0], err
}
