package domain

import (
	"time"
)

type Code struct {
	Id       int64     `json:"id" db:"id"`
	Code     string    `json:"code" db:"code"`
	Exp      time.Time `json:"exp" db:"exp"`
	UserId   int64     `json:"user_id" db:"user_id"`
	ClientId int64     `json:"client_id" db:"client_id"`
}

type ExchangeCode struct {
	Code string `json:"code" binding:"required"`
}
