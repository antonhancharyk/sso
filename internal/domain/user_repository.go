package domain

type UserRepo interface {
	GetUserByID(id int64) (User, error)
	GetUserByEmail(email string) (User, error)
	CreateUser(user User) error
}
