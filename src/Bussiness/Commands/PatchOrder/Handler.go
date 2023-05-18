package PatchOrder

import (
	"4_zad/Models"
	"4_zad/Services"
	"context"
	"golang.org/x/exp/slices"
)

type Handler struct {
	basketRepository Services.IRepository[Models.Basket]
}

func NewHandler(basketRepository Services.IRepository[Models.Basket]) *Handler {
	return &Handler{basketRepository: basketRepository}
}

func (g Handler) Handle(ctx context.Context, request *Request) (*Response, error) {
	basket := g.basketRepository.GetById(request.BasketId)
	orderIdx := slices.IndexFunc(basket.Orders, func(order Models.Order) bool { return order.ID == request.OrderId })
	if *request.OrderPatch.Quantity < 0 {
		basket.Orders[orderIdx].Quantity -= (uint)(-(*request.OrderPatch.Quantity))
	} else {
		basket.Orders[orderIdx].Quantity += (uint)(*request.OrderPatch.Quantity)
	}
	g.basketRepository.Update(basket)
	return NewResponse(basket.Orders[orderIdx]), nil //TODO: instead of nil, should probably return true error
}
