package routers

import (
	"database/sql"

	v1 "github.com/dellta103/s3-golang-assignment/internal/handler/rest/v1"
	"github.com/dellta103/s3-golang-assignment/internal/repository"
	"github.com/dellta103/s3-golang-assignment/internal/service/user"
	"github.com/go-chi/chi/v5"
)

func InitRoute(db *sql.DB) *chi.Mux {
	repo := repository.NewRepo(db)
	userServ := user.New(repo)
	handler := v1.NewHandler(userServ)

	r := chi.NewRouter()
	r.Route("/api/v1", func(api chi.Router) {
		r.Route("/users", userRouter(handler))
	})

	return r
}

func userRouter(h v1.Handler) func(api chi.Router) {
	return func(api chi.Router) {
	}
}
