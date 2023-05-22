package GetUser

import "4_zad/Models"

type Response struct {
	User Models.User
}

func NewResponse(user Models.User) *Response {
	return &Response{User: user}
}
