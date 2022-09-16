package repository

import (
	"database/sql"

	"github.com/dellta103/s3-golang-assignment/internal/repository/user"
)

type IRepo interface {
	User() user.IUser
}

type impl struct {
	db   *sql.DB
	user user.IUser
}

func NewRepo(db *sql.DB) IRepo {
	return impl{
		db:   db,
		user: user.New(db),
	}
}

func (i impl) User() user.IUser {
	return i.user
}
