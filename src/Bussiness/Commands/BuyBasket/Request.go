package BuyBasket

type Request struct {
	BasketId uint
	Orders   []struct {
		ProductId uint
		Quantity  uint
	}
}

func NewRequest(basketId uint, orders []struct {
	ProductId uint
	Quantity  uint
}) *Request {
	return &Request{BasketId: basketId, Orders: orders}
}
