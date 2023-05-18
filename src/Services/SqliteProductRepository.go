package Services

import (
	"4_zad/Models"
	"gorm.io/gorm"
	"time"
)

type SqliteProductRepository struct {
	db *gorm.DB
}

func NewSqliteProductRepository(db *gorm.DB) IRepository[Models.Product] {
	migrationError := db.AutoMigrate(&Models.Product{})
	if migrationError != nil {
		panic(migrationError)
	}

	return &SqliteProductRepository{db}
}

func (repo *SqliteProductRepository) GetAll() []Models.Product {
	var products []Models.Product
	repo.db.Find(&products)
	return products
}

func (repo *SqliteProductRepository) GetById(id uint) Models.Product {
	var product Models.Product
	repo.db.First(&product, id)
	return product
}

func (repo *SqliteProductRepository) GetWithFilter() []Models.Product {
	panic("Not implemented!")
}

func (repo *SqliteProductRepository) Add(product Models.Product) Models.Product {
	var productToCreate Models.Product
	productToCreate = product
	repo.db.Create(&productToCreate)
	return product
}

func (repo *SqliteProductRepository) Update(product Models.Product) Models.Product {
	repo.db.Save(&product)
	return product
}

func (repo *SqliteProductRepository) Delete(id uint) Models.Product {
	repo.db.Delete(&Models.Product{}, id)
	return Models.Product{}
}

func (repo *SqliteProductRepository) Patch(id uint, productPatch any) Models.Product {
	var productToUpdate Models.Product
	repo.db.First(&productToUpdate, id)
	if productPatch.(Models.PatchProductBinding).Name != nil {
		productToUpdate.Name = *productPatch.(Models.PatchProductBinding).Name
	}
	if productPatch.(Models.PatchProductBinding).Price != nil {
		productToUpdate.Price = *productPatch.(Models.PatchProductBinding).Price
	}
	if productPatch.(Models.PatchProductBinding).Category != nil {
		productToUpdate.Category = *productPatch.(Models.PatchProductBinding).Category
	}
	if productPatch.(Models.PatchProductBinding).ImageURL != nil {
		productToUpdate.ImageURL = *productPatch.(Models.PatchProductBinding).ImageURL
	}
	if productPatch.(Models.PatchProductBinding).Quantity != nil {
		productToUpdate.Quantity = *productPatch.(Models.PatchProductBinding).Quantity
	}

	productToUpdate.UpdatedAt = time.Now()
	repo.db.Save(&productToUpdate)
	return productToUpdate
}
