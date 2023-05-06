package GetProducts

import "4_zad/Models"

type Response struct {
	Products []Models.Product
}

func NewResponse(products []Models.Product) *Response {
	return &Response{Products: products}
}
