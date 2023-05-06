package GetProduct

import "4_zad/Models"

type Response struct {
	Product Models.Product
}

func NewResponse(product Models.Product) *Response {
	return &Response{Product: product}
}
