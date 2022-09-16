package user

import (
	"context"
	"database/sql"

	"github.com/dellta103/s3-golang-assignment/internal/model"
)

type IUser interface {
	// CreateUser creates a new user
	CreateUser(ctx context.Context, user model.User) (model.User, error)
}

type impl struct {
	db *sql.DB
}

func New(db *sql.DB) IUser {
	return impl{
		db: db,
	}
}
