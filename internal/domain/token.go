package domain

type Token struct {
	Id           int64  `json:"id" db:"id"`
	UserId       int64  `json:"user_id" db:"user_id"`
	AccessToken  string `json:"access_token" db:"access_token"`
	RefreshToken string `json:"refresh_token" db:"refresh_token"`
}

type CreateToken struct {
	UserId       int64  `json:"user_id" db:"user_id"`
	AccessToken  string `json:"access_token" db:"access_token"`
	RefreshToken string `json:"refresh_token" db:"refresh_token"`
}

type ValidateToken struct {
	Token string `form:"token" binding:"required"`
}
