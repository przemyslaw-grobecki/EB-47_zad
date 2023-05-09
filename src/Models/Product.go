package Models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name     string
	Category string
	Price    uint
	Quantity uint
	ImageURL string
}
