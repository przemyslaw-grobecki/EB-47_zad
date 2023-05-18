package GetBaskets

import "4_zad/Models"

type Response struct {
	Baskets []Models.Basket
}

func NewResponse(baskets []Models.Basket) *Response {
	return &Response{Baskets: baskets}
}
