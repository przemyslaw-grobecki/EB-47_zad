package PatchBasket

import "4_zad/Models"

type Response struct {
	Basket Models.Basket
}

func NewResponse(basket Models.Basket) *Response {
	return &Response{Basket: basket}
}
