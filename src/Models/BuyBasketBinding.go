package Models

type BuyBasketBinding struct {
	Orders []struct {
		ProductId uint
		Quantity  uint
	}
}
