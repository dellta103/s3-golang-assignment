package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

// ConnectDB
func ConnectDB(dbURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}
	return db, nil
}
