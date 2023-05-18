package AddOrder

type Request struct {
	BasketId  uint
	ProductId uint
	Quantity  uint
}

func NewRequest(basketId uint, productId uint, quantity uint) *Request {
	return &Request{BasketId: basketId, ProductId: productId, Quantity: quantity}
}
