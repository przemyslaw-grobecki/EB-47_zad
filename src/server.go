package main

import (
	"4_zad/Controllers"
	"4_zad/Services"
	"github.com/labstack/echo/v4"
)

func main() {
	//Controllers
	e := echo.New()
	repository := Services.NewProductRepository()
	Controllers.BootstrapProductsController(e, repository)

	e.Logger.Fatal(e.Start(":1323"))
}
