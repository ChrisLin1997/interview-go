package db

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var db *sqlx.DB

func Connect() {
	dbHost := os.Getenv("POSTGRES_HOST")
	dbPort := os.Getenv("POSTGRES_PORT")
	dbUser := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB_NAME")
	dbSslMode := os.Getenv("POSTGRES_SSL_MODE")

	sourceName := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", dbHost, dbPort, dbUser, dbPassword, dbName, dbSslMode)
	fmt.Println(sourceName)

	postgre, err := sqlx.Open("postgres", sourceName)

	if err != nil {
		log.Fatalln(err)
		defer db.Close()
	}

	db = postgre
}

func Get() *sqlx.DB {
	return db
}
