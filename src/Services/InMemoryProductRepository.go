package Services

import (
	"4_zad/Models"
	"gorm.io/gorm"
	"time"
)

//
//{
//"products": [
//{
//"name": "Potato",
//"category": "Food",
//"price": 2.50,
//"quantity": 50,
//"imageURL": "https://images.pexels.com/photos/144248/potatoes-vegetables-erdfrucht-bio-144248.jpeg"
//},
//{
//"name": "Tomato",
//"category": "Food",
//"price": 3.60,
//"quantity": 70,
//"imageURL": "https://images.pexels.com/photos/96616/pexels-photo-96616.jpeg?auto=compress&cs=tinysrgb&w=1260&h=750&dpr=1"
//},
//{
//"name": "Strawberry",
//"category": "Food",
//"price": 4.5,
//"quantity": 160,
//"imageURL": "https://images.pexels.com/photos/1788912/pexels-photo-1788912.jpeg?auto=compress&cs=tinysrgb&w=1260&h=750&dpr=1"
//},
//{
//"name": "Raspberry",
//"category": "Food",
//"price": 5.5,
//"quantity": 120,
//"imageURL": "https://images.pexels.com/photos/3429782/pexels-photo-3429782.jpeg?auto=compress&cs=tinysrgb&w=1260&h=750&dpr=1"
//},
//{
//"name": "Watermelon",
//"category": "Food",
//"price": 10,
//"quantity": 11,
//"imageURL": "https://images.pexels.com/photos/5946081/pexels-photo-5946081.jpeg?auto=compress&cs=tinysrgb&w=1260&h=750&dpr=1"
//}
//]
//}

type InMemoryProductRepository struct {
}

var products = []Models.Product{
	{
		Model: gorm.Model{
			ID:        0,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			DeletedAt: gorm.DeletedAt{},
		},
		Name: "watermelon",
	},
	{
		Model: gorm.Model{
			ID:        1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			DeletedAt: gorm.DeletedAt{},
		},
		Name: "apple",
	},
	{
		Model: gorm.Model{
			ID:        2,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			DeletedAt: gorm.DeletedAt{},
		},
		Name: "spinach",
	},
}

func NewInMemoryRepository() IRepository[Models.Product] {
	return &InMemoryProductRepository{}
}

func (repo *InMemoryProductRepository) GetAll() []Models.Product {
	return products
}

func (repo *InMemoryProductRepository) GetById(id uint) Models.Product {
	for _, product := range products {
		if product.ID == id {
			return product
		}
	}
	panic("No product with given Id found.")
}

func (repo *InMemoryProductRepository) GetWithFilter() []Models.Product {
	panic("Not implemented.")
}

func (repo *InMemoryProductRepository) Add(product Models.Product) Models.Product {
	newProduct := Models.Product{
		Name: product.Name,
		Model: gorm.Model{
			ID:        product.ID,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			DeletedAt: gorm.DeletedAt{},
		},
	}
	products = append(products, newProduct)
	return newProduct
}

func (repo *InMemoryProductRepository) Delete(id uint) Models.Product {
	var indexToDelete int
	var productFound = false
	var productToDelete Models.Product
	for index, product := range products {
		if product.ID == id {
			indexToDelete = index
			productToDelete = product
			productFound = true
		}
	}
	if !productFound {
		panic("No product with given Id found.")
	}
	products = append(products[:indexToDelete], products[indexToDelete+1:]...)
	return productToDelete
}

func (repo *InMemoryProductRepository) Patch(id uint, productPatch any) Models.Product {
	productToPatch := repo.Delete(id)
	productToPatch.Model.UpdatedAt = time.Now()
	if productPatch.(Models.Product).Name != "" {
		productToPatch.Name = productPatch.(Models.Product).Name
	}
	products = append(products, productToPatch)
	return productToPatch
}
