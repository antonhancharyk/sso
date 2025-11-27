package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/antongoncharik/sso/internal/infra/db"
	"github.com/antongoncharik/sso/internal/infra/logger"
	"github.com/antongoncharik/sso/internal/infra/security"
	"github.com/antongoncharik/sso/internal/repository"
	"github.com/antongoncharik/sso/internal/service"
	"github.com/antongoncharik/sso/internal/transport/httpserver"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	log := logger.New()

	keys, err := security.MustLoad()
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}

	err = db.Connect()
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}
	defer db.Close()

	repos := repository.NewRepositories(db.Get())
	services := service.NewServices(service.ServiceDeps{
		UserRepo:   repos.User,
		ClientRepo: repos.Client,
		CodeRepo:   repos.Code,
		TokenRepo:  repos.Token,
		RSA:        keys,
	})
	hdl := httpserver.NewHandler(services)
	r := httpserver.GetRoutes(hdl)

	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	go func() {
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Error(fmt.Errorf("could not listen on :8080: %v", err))
			os.Exit(1)
		}
	}()
	log.Info("server started on :8080")

	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	log.Info("shutting down server...")
	err = server.Shutdown(ctx)
	if err != nil {
		log.Error(fmt.Errorf("server forced to shutdown: %v", err))
	}
	log.Info("server exiting")
}
