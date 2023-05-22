package LoginUser

import (
	"4_zad/Models"
	"4_zad/Services"
	"context"
	"errors"
	"golang.org/x/exp/slices"
)

type Handler struct {
	userRepository Services.IRepository[Models.User]
}

func NewHandler(userRepository Services.IRepository[Models.User]) *Handler {
	return &Handler{userRepository: userRepository}
}

func (g Handler) Handle(ctx context.Context, request *Request) (*Response, error) {
	users := g.userRepository.GetAll()
	index := slices.IndexFunc(users, func(user Models.User) bool {
		return user.Email == request.Email
	})
	if index == -1 {
		return nil, errors.New("No Account Found")
	}
	return NewResponse(users[index]), nil //TODO: instead of nil, should probably return true error
}
