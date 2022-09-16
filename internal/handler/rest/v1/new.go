package v1

import "github.com/dellta103/s3-golang-assignment/internal/service/user"

type Handler struct {
	userServ user.UserServ
}

func NewHandler(userServ user.UserServ) Handler {
	return Handler{userServ: userServ}
}
