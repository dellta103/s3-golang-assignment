package v1

import "net/http"

type Error struct {
	Status int    `json:"status"`
	Code   string `json:"code"`
	Desc   string `json:"desc"`
}

var (
	ErrNameCannotBeBlank = Error{Status: http.StatusBadRequest, Code: "invalid_input", Desc: "name cannot be blank"}
)
