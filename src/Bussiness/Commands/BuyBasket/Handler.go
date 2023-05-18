package BuyBasket

import (
	"4_zad/Models"
	"4_zad/Services"
	"context"
	"errors"
)

func filter[T Models.Product](ss []T, test func(T) bool) (ret []T) {
	for _, s := range ss {
		if test(s) {
			ret = append(ret, s)
		}
	}
	return
}

type Handler struct {
	basketRepository  Services.IRepository[Models.Basket]
	productRepository Services.IRepository[Models.Product]
}

func NewHandler(basketRepository Services.IRepository[Models.Basket],
	productRepository Services.IRepository[Models.Product]) *Handler {
	return &Handler{basketRepository: basketRepository, productRepository: productRepository}
}

func (g Handler) Handle(ctx context.Context, request *Request) (*Response, error) {
	allProducts := g.productRepository.GetAll()
	var productsToBuy []Models.Product
	for _, order := range request.Orders {
		for _, product := range allProducts {
			if order.ProductId == product.ID {
				if product.Quantity < order.Quantity {
					return nil, errors.New("One of products out of stock.")
				}
				product.Quantity -= order.Quantity
				productsToBuy = append(productsToBuy, product)
			}
		}
	}
	for _, productToBuy := range productsToBuy {
		g.productRepository.Patch(productToBuy.ID, Models.PatchProductBinding{
			Name:     nil,
			Category: nil,
			Price:    nil,
			Quantity: &productToBuy.Quantity,
			ImageURL: nil,
		})
	}

	return NewResponse(), nil //TODO: instead of nil, should probably return true error
}
