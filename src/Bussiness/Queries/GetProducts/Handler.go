package GetProducts

import (
	"4_zad/Models"
	"4_zad/Services"
	"context"
)

type Handler struct {
	productRepository Services.IRepository[Models.Product]
}

func NewHandler(productRepository Services.IRepository[Models.Product]) *Handler {
	return &Handler{productRepository: productRepository}
}

func (g Handler) Handle(ctx context.Context, request *Request) (*Response, error) {
	return &Response{Products: g.productRepository.GetAll()}, nil //TODO: instead of nil, should probably return true error
}
