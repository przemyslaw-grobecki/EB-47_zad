package Models

type NewOrderBinding struct {
	ProductId uint `json:"ProductId" binding:"required"`
	Quantity  uint `json:"Quantity" binding:"required"`
}
