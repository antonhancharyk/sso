package domain

type User struct {
	Id           int64  `json:"id" db:"id"`
	Email        string `json:"email" db:"email"`
	PasswordHash string `json:"password_hash" db:"password_hash"`
}

type LoginRegisterForm struct {
	App         string `form:"app" binding:"required"`
	RedirectUri string `form:"redirect_uri" binding:"required"`
}

type Register struct {
	Email       string `form:"email" binding:"required,email"`
	Password    string `form:"password" binding:"required"`
	App         string `form:"app" binding:"required"`
	RedirectUri string `form:"redirect_uri" binding:"required"`
}

type Login struct {
	Email       string `form:"email" binding:"required,email"`
	Password    string `form:"password" binding:"required"`
	App         string `form:"app" binding:"required"`
	RedirectUri string `form:"redirect_uri" binding:"required"`
}
