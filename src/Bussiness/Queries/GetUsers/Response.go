package GetUsers

import "4_zad/Models"

type Response struct {
	Users []Models.User
}

func NewResponse(users []Models.User) *Response {
	return &Response{Users: users}
}
