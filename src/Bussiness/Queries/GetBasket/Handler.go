package GetBasket

import (
	"4_zad/Models"
	"4_zad/Services"
	"context"
)

type Handler struct {
	basketRepository Services.IRepository[Models.Basket]
}

func NewHandler(basketRepository Services.IRepository[Models.Basket]) *Handler {
	return &Handler{basketRepository: basketRepository}
}

func (g Handler) Handle(ctx context.Context, request *Request) (*Response, error) {
	return &Response{Basket: g.basketRepository.GetById(request.Id)}, nil
}
