package Models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	ProductId uint
	Quantity  uint
	BasketId  uint
}
