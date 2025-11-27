package domain

type TokenRepo interface {
	CreateToken(createToken CreateToken) error
	GetToken(token string) (Token, error)
}
