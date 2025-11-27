package domain

type Client struct {
	Id          int64  `json:"id" db:"id"`
	App         string `json:"app" db:"app"`
	RedirectURI string `json:"redirect_uri" db:"redirect_uri"`
}
