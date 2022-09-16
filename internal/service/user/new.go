package user

import (
	"github.com/dellta103/s3-golang-assignment/internal/repository"
)

type UserServ interface {
}

type impl struct {
	repo repository.IRepo
}

func New(repo repository.IRepo) UserServ {
	return impl{repo: repo}
}
