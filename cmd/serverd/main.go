package main

import (
	"log"
	"net/http"
	"os"

	"github.com/dellta103/s3-golang-assignment/cmd/serverd/routers"
	"github.com/dellta103/s3-golang-assignment/pkg/db"
)

func main() {
	// DB connection
	db, err := db.ConnectDB(os.Getenv("DB_URL"))
	if err != nil {
		log.Fatal(err)
	}

	// TODO: init router
	// get: /v1/users
	r := routers.InitRoute(db)

	// TODO: start server
	if err := http.ListenAndServe(":5000", r); err != nil {
		log.Fatal(err)
	}
}
