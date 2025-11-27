package domain

type ClientRepo interface {
	GetClient(app string) (Client, error)
}
