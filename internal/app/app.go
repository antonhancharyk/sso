package app

import (
	"log"

	"github.com/antongoncharik/sso/internal/api/http"
	"github.com/antongoncharik/sso/internal/api/http/handler"
	"github.com/antongoncharik/sso/internal/config"
	"github.com/antongoncharik/sso/internal/database"
	"github.com/antongoncharik/sso/internal/repository"
	"github.com/antongoncharik/sso/internal/service"
)

func Run() {
	keys, err := config.MustLoad()
	if err != nil {
		log.Fatal(err)
	}

	database.Connect()
	defer database.Close()

	db := database.Get()

	repo := repository.New(db)
	svc := service.New(repo, keys)
	hdl := handler.New(svc)

	http.RunHTTP(hdl)
}
