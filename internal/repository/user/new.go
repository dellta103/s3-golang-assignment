package user

import (
	"database/sql"
)

type IUser interface {
}

type impl struct {
	db *sql.DB
}

func New(db *sql.DB) IUser {
	return impl{
		db: db,
	}
}
