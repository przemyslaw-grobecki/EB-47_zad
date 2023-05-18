package AddBasket

import (
	"4_zad/Models"
	"4_zad/Services"
	"context"
	"gorm.io/gorm"
)

type Handler struct {
	basketRepository Services.IRepository[Models.Basket]
}

func NewHandler(basketRepository Services.IRepository[Models.Basket]) *Handler {
	return &Handler{basketRepository: basketRepository}
}

func (g Handler) Handle(ctx context.Context, request *Request) (*Response, error) {
	return NewResponse(g.basketRepository.Add(Models.Basket{Status: "open", Model: gorm.Model{ID: request.Id}})), nil //TODO: instead of nil, should probably return true error
}
