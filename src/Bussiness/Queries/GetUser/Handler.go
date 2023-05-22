package GetUser

import (
	"4_zad/Models"
	"4_zad/Services"
	"context"
)

type Handler struct {
	userRepository Services.IRepository[Models.User]
}

func NewHandler(userRepository Services.IRepository[Models.User]) *Handler {
	return &Handler{userRepository: userRepository}
}

func (g Handler) Handle(ctx context.Context, request *Request) (*Response, error) {
	return &Response{User: g.userRepository.GetById(request.Id)}, nil
}
