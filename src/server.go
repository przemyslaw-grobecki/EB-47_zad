package main

import (
	"4_zad/Controllers"
	"4_zad/Services"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.DefaultCORSConfig))
	db, dbError := gorm.Open(sqlite.Open("sqlite.db"), &gorm.Config{})
	if dbError != nil {
		panic(dbError)
	}

	//repository := Services.NewInMemoryRepository() //LEGACY: used as an initial data storage mock
	productRepository := Services.NewSqliteProductRepository(db)
	basketRepository := Services.NewSqliteBasketRepository(db)
	userRepository := Services.NewSqliteUserRepository(db)

	Controllers.BootstrapProductsController(e, productRepository)
	Controllers.BootstrapBasketsController(e, basketRepository, productRepository)
	Controllers.BootstrapUsersController(e, userRepository)

	e.Logger.Fatal(e.Start(":1323"))
}
