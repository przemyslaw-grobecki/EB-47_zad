package DeleteProduct

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
	return NewResponse(g.productRepository.Delete(request.Id)), nil //TODO: instead of nil, should probably return true error
}
