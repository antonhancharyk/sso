package entity

import (
	"crypto/rsa"
	"time"
)

type User struct {
	Id           int64  `json:"id" db:"id"`
	Email        string `json:"email" db:"email"`
	PasswordHash string `json:"password_hash" db:"password_hash"`
}

type Client struct {
	Id          int64  `json:"id" db:"id"`
	App         string `json:"app" db:"app"`
	RedirectURI string `json:"redirect_uri" db:"redirect_uri"`
}

type Token struct {
	Id           int64  `json:"id" db:"id"`
	UserId       int64  `json:"user_id" db:"user_id"`
	AccessToken  string `json:"access_token" db:"access_token"`
	RefreshToken string `json:"refresh_token" db:"refresh_token"`
}

type Code struct {
	Id       int64     `json:"id" db:"id"`
	Code     string    `json:"code" db:"code"`
	Exp      time.Time `json:"exp" db:"exp"`
	UserId   int64     `json:"user_id" db:"user_id"`
	ClientId int64     `json:"client_id" db:"client_id"`
}

type LoginRegisterForm struct {
	App         string `form:"app" binding:"required"`
	RedirectUri string `form:"redirect_uri" binding:"required"`
}

type Register struct {
	Email       string `form:"email" binding:"required,email"`
	Password    string `form:"password" binding:"required"`
	App         string `form:"app" binding:"required"`
	RedirectURI string `form:"redirect_uri" binding:"required"`
}

type Login struct {
	Email       string `form:"email" binding:"required,email"`
	Password    string `form:"password" binding:"required"`
	App         string `form:"app" binding:"required"`
	RedirectURI string `form:"redirect_uri" binding:"required"`
}

type ValidateToken struct {
	Token string `form:"token" binding:"required"`
}

type ExchangeCode struct {
	Code string `json:"code" binding:"required"`
}

type CreateToken struct {
	UserId       int64  `json:"user_id" db:"user_id"`
	AccessToken  string `json:"access_token" db:"access_token"`
	RefreshToken string `json:"refresh_token" db:"refresh_token"`
}

type RSAKeys struct {
	PrivateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
}
