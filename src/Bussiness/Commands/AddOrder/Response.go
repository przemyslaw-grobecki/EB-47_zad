package AddOrder

import "4_zad/Models"

type Response struct {
	Order Models.Order
}

func NewResponse(order Models.Order) *Response {
	return &Response{Order: order}
}
