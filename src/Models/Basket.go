package Models

import "gorm.io/gorm"

type Basket struct {
	gorm.Model
	Status string
	Orders []Order `gorm:"constraint:OnUpdate:CASCADE;" gorm:"foreignkey:BasketId"`
}
