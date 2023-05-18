package Services

import (
	"4_zad/Models"
	"gorm.io/gorm"
	"time"
)

type SqliteBasketRepository struct {
	db *gorm.DB
}

func NewSqliteBasketRepository(db *gorm.DB) IRepository[Models.Basket] {
	basketMigrationError := db.AutoMigrate(&Models.Basket{})
	if basketMigrationError != nil {
		panic(basketMigrationError)
	}

	orderMigrationError := db.AutoMigrate(&Models.Order{})
	if orderMigrationError != nil {
		panic(orderMigrationError)
	}

	return &SqliteBasketRepository{db}
}

func (repo *SqliteBasketRepository) GetAll() []Models.Basket {
	var baskets []Models.Basket
	repo.db.Preload("Orders").Find(&baskets)
	return baskets
}

func (repo *SqliteBasketRepository) GetById(id uint) Models.Basket {
	var basket Models.Basket
	repo.db.Preload("Orders").First(&basket, id)
	return basket
}

func (repo *SqliteBasketRepository) GetWithFilter() []Models.Basket {
	panic("Not implemented!")
}

func (repo *SqliteBasketRepository) Add(basket Models.Basket) Models.Basket {
	repo.db.Create(&basket)
	return basket
}

func (repo *SqliteBasketRepository) Update(basket Models.Basket) Models.Basket {
	repo.db.Save(&basket)
	return basket
}

func (repo *SqliteBasketRepository) Delete(id uint) Models.Basket {
	repo.db.Delete(&Models.Basket{}, id)
	return Models.Basket{}
}

func (repo *SqliteBasketRepository) Patch(id uint, basketPatch any) Models.Basket {
	var basketToUpdate Models.Basket
	repo.db.First(&basketToUpdate, id)
	if basketPatch.(Models.PatchBasketBinding).Status != nil {
		basketToUpdate.Status = *basketPatch.(Models.PatchBasketBinding).Status
	}
	basketToUpdate.UpdatedAt = time.Now()
	repo.db.Save(&basketToUpdate)
	return basketToUpdate
}
