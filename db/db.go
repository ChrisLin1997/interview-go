package db

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var db *sqlx.DB

func Connect() {
	postgre, err := sqlx.Open("postgres", "host=127.0.0.1 port=5432 user=postgres password=0000 dbname=postgres sslmode=disable")

	if err != nil {
		log.Fatalln(err)
		defer db.Close()
	}

	db = postgre
}

func Get() *sqlx.DB {
	return db
}
