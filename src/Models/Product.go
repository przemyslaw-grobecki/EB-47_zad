package Models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name     string
	Category string
	Price    float64
	Quantity uint
	ImageURL string
}
