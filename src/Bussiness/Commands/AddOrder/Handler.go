package AddOrder

import (
	"4_zad/Models"
	"4_zad/Services"
	"context"
	"errors"
	"golang.org/x/exp/slices"
)

type Handler struct {
	basketRepository  Services.IRepository[Models.Basket]
	productRepository Services.IRepository[Models.Product]
}

func NewHandler(basketRepository Services.IRepository[Models.Basket],
	productRepository Services.IRepository[Models.Product]) *Handler {
	return &Handler{basketRepository: basketRepository, productRepository: productRepository}
}

func (g Handler) Handle(ctx context.Context, request *Request) (*Response, error) {
	basket := g.basketRepository.GetById(request.BasketId)
	product := g.productRepository.GetById(request.ProductId)

	if product.Quantity < request.Quantity {
		return nil, errors.New("ordered product is not available")
	}
	if slices.ContainsFunc(basket.Orders, func(order Models.Order) bool { return order.ProductId == request.ProductId }) {
		orderToUpdateIndex := slices.IndexFunc(basket.Orders, func(order Models.Order) bool { return order.ProductId == request.ProductId })
		basket.Orders[orderToUpdateIndex].Quantity += request.Quantity
		g.basketRepository.Update(basket)
		return NewResponse(basket.Orders[orderToUpdateIndex]), nil
	}
	order := Models.Order{BasketId: request.BasketId, ProductId: request.ProductId, Quantity: request.Quantity}
	basket.Orders = append(basket.Orders, order)
	g.basketRepository.Update(basket)
	return NewResponse(order), nil //TODO: instead of nil, should probably return true error
}
