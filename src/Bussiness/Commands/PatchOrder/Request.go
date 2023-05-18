package PatchOrder

import "4_zad/Models"

type Request struct {
	BasketId   uint
	OrderId    uint
	OrderPatch Models.PatchOrderBinding
}

func NewRequest(basketId uint, orderId uint, orderPatchBinding Models.PatchOrderBinding) *Request {
	return &Request{BasketId: basketId, OrderId: orderId, OrderPatch: orderPatchBinding}
}
