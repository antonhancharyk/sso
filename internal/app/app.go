package app

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	httpApp "github.com/antongoncharik/sso/internal/api/http"
	"github.com/antongoncharik/sso/internal/api/http/handler"
	"github.com/antongoncharik/sso/internal/config"
	"github.com/antongoncharik/sso/internal/database"
	"github.com/antongoncharik/sso/internal/repository"
	"github.com/antongoncharik/sso/internal/service"
	"github.com/antongoncharik/sso/pkg/logger"
)

func Run() {
	log := logger.New()

	keys, err := config.MustLoad()
	if err != nil {
		log.Fatal(err)
	}

	database.Connect(log)
	defer database.Close()

	db := database.Get()

	repo := repository.New(db)
	svc := service.New(repo, keys)
	hdl := handler.New(svc)

	r := httpApp.GetRoutes(hdl)

	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	go func() {
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatal(errors.New(fmt.Sprintf("could not listen on :8080: %v\n", err)))
		}
	}()
	log.Info("server started on :8080")

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	log.Info("shutting down server...")
	err = server.Shutdown(ctx)
	if err != nil {
		log.Fatal(errors.New(fmt.Sprintf("server forced to shutdown: %v\n", err)))
	}
	log.Info("server exiting")
}
