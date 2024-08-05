package database

import (
	"errors"
	"fmt"
	"os"

	"github.com/antongoncharik/sso/pkg/logger"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var db *sqlx.DB

func Connect(log *logger.Logger) {
	connectStr := fmt.Sprintf("user=%s dbname=%s sslmode=disable password=%s host=db", os.Getenv("DB_USER"), os.Getenv("DB_NAME"), os.Getenv("DB_PASSWORD"))

	conn, err := sqlx.Connect("postgres", connectStr)
	if err != nil {
		log.Fatal(err)
	}
	err = conn.Ping()
	if err != nil {
		log.Fatal(err)
	}

	driver, err := postgres.WithInstance(conn.DB, &postgres.Config{})
	if err != nil {
		log.Fatal("could not create migrate instance: %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		os.Getenv("DB_NAME"), driver)
	if err != nil {
		log.Fatal(errors.New(fmt.Sprintf("could not create migrate instance: %v", err)))
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		log.Fatal(errors.New(fmt.Sprintf("could not run up migrations: %v", err)))
	}

	log.Info("migrations applied successfully!")

	db = conn
}

func Get() *sqlx.DB {
	return db
}

func Close() {
	db.Close()
}
