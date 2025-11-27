package domain

type CodeRepo interface {
	CreateCode(code Code) error
	GetCode(exchangeCode ExchangeCode) (Code, error)
}
