package RegisterUser

import (
	"4_zad/Models"
	"4_zad/Services"
	"context"
	"golang.org/x/exp/rand"
	"gorm.io/gorm"
)

type Handler struct {
	userRepository Services.IRepository[Models.User]
}

func NewHandler(userRepository Services.IRepository[Models.User]) *Handler {
	return &Handler{userRepository: userRepository}
}

func (g Handler) Handle(ctx context.Context, request *Request) (*Response, error) {
	return NewResponse(g.userRepository.Add(Models.User{
		Model:    gorm.Model{ID: uint(rand.Uint32())},
		Email:    request.Email,
		Password: request.Password,
	})), nil //TODO: instead of nil, should probably return true error
}
