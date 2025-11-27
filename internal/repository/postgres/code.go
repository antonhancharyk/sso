package postgres

import (
	"errors"

	"github.com/antongoncharik/sso/internal/domain"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type CodeRepo struct {
	db *sqlx.DB
}

func NewCodeRepo(db *sqlx.DB) domain.CodeRepo {
	return &CodeRepo{db: db}
}

func (r *CodeRepo) CreateCode(code domain.Code) error {
	_, err := r.db.Exec(`INSERT INTO codes (code, exp, user_id, client_id)
	VALUES ($1, $2, $3, $4)`, code.Code, code.Exp, code.UserId, code.ClientId)

	return err
}

func (r *CodeRepo) GetCode(exchangeCode domain.ExchangeCode) (domain.Code, error) {
	codeData := []domain.Code{}

	err := r.db.Select(&codeData, "SELECT id, user_id, exp FROM codes WHERE code = $1", exchangeCode.Code)

	if len(codeData) == 0 {
		return domain.Code{}, errors.New("code not found")
	}

	return codeData[0], err
}
