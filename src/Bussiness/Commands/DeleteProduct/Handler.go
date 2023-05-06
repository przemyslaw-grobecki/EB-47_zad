package DeleteProduct

import (
	"4_zad/Services"
	"context"
)

type Handler struct {
	productRepository *Services.ProductRepository
}

func NewHandler(productRepository *Services.ProductRepository) *Handler {
	return &Handler{productRepository: productRepository}
}

func (g Handler) Handle(ctx context.Context, request *Request) (*Response, error) {
	return NewResponse(g.productRepository.DeleteProduct(request.Id)), nil //TODO: instead of nil, should probably return true error
}
