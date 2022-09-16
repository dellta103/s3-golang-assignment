package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/dellta103/s3-golang-assignment/pkg/db"
	"github.com/go-chi/chi/v5"
)

func main() {

	// DB connection
	db, err := db.ConnectDB(os.Getenv("DB_URL"))
	if err != nil {
		log.Fatal(err)
	}

	// TODO: init router
	// get: /v1/users
	r := initRouter(db)

	// TODO: start server
	if err := http.ListenAndServe(":5000", r); err != nil {
		log.Fatal(err)
	}
}

func initRouter(db *sql.DB) *chi.Mux {
	r := chi.NewRouter()
	return r
}
