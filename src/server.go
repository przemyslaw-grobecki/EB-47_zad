package main

import (
	"4_zad/Controllers"
	"4_zad/Services"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	//db, dbError := gorm.Open(sqlite.Open("sqlite.db"), &gorm.Config{})
	//if dbError != nil {
	//	panic(dbError)
	//}
	//product := Models.Product{
	//	Name: "pig",
	//}
	//result := db.Create(&product)
	//if result != nil {
	//
	//}
	//
	//foundProd := db.Find(&product)
	//
	//if foundProd != nil {
	//
	//}

	repository := Services.NewInMemoryRepository() //LEGACY: used as an initial data storage mock
	//repository := Services.NewSqliteProductRepository(db)

	Controllers.BootstrapProductsController(e, repository)

	e.Logger.Fatal(e.Start(":1323"))
}
