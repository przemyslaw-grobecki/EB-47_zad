package Services

import (
	"4_zad/Models"
	"gorm.io/gorm"
)

type SqliteProductRepository struct {
	db *gorm.DB
}

func NewSqliteProductRepository(db *gorm.DB) *SqliteProductRepository {
	migrationError := db.AutoMigrate(&Models.Product{})
	if migrationError != nil {
		panic(migrationError)
	}

	return &SqliteProductRepository{db}
}

func (repo *SqliteProductRepository) GetAllProducts() []Models.Product {
	panic("Not implemented!")
}

func (repo *SqliteProductRepository) GetProductById(id uint) Models.Product {
	panic("Not implemented!")
}

func (repo *SqliteProductRepository) GetProductsWithFilter() {
	panic("Not implemented!")
}

func (repo *SqliteProductRepository) AddProduct(id uint, name string) Models.Product {
	panic("Not implemented!")
}

func (repo *SqliteProductRepository) DeleteProduct(id uint) Models.Product {
	panic("Not implemented!")
}

func (repo *SqliteProductRepository) PatchProduct(id uint, productPatch Models.PatchProductBinding) Models.Product {
	panic("Not implemented!")
}
