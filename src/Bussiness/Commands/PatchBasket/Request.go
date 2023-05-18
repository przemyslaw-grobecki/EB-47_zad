package PatchBasket

import (
	"4_zad/Models"
)

type Request struct {
	Id          uint
	BasketPatch Models.PatchBasketBinding
}

func NewRequest(id uint, basketPatch Models.PatchBasketBinding) *Request {
	return &Request{Id: id, BasketPatch: basketPatch}
}
