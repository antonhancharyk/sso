package database

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var db *sqlx.DB

func Connect() {
	connectStr := fmt.Sprintf("user=%s dbname=%s sslmode=disable password=%s host=db", os.Getenv("DB_USER"), os.Getenv("DB_NAME"), os.Getenv("DB_PASSWORD"))

	conn, err := sqlx.Connect("postgres", connectStr)
	if err != nil {
		log.Fatalln(err)
	}
	err = conn.Ping()
	if err != nil {
		log.Fatalln(err)
	}

	db = conn
}

func Get() *sqlx.DB {
	return db
}

func Close() {
	db.Close()
}
