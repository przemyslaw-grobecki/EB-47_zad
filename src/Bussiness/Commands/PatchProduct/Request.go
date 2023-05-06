package PatchProduct

import (
	"4_zad/Models"
)

type Request struct {
	Id           uint
	ProductPatch Models.PatchProductBinding
}

func NewRequest(id uint, productPatch Models.PatchProductBinding) *Request {
	return &Request{Id: id, ProductPatch: productPatch}
}
