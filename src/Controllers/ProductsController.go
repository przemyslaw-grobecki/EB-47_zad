package Controllers

import (
	"4_zad/Bussiness/Commands/AddProduct"
	"4_zad/Bussiness/Commands/DeleteProduct"
	"4_zad/Bussiness/Commands/PatchProduct"
	"4_zad/Bussiness/Queries/GetProduct"
	"4_zad/Bussiness/Queries/GetProducts"
	"4_zad/Models"
	"4_zad/Services"
	"github.com/labstack/echo/v4"
	"github.com/mehdihadeli/go-mediatr"
	"net/http"
	"strconv"
)

func BootstrapProductsController(e *echo.Echo, productRepository *Services.ProductRepository) {
	_ = mediatr.RegisterRequestHandler[*GetProduct.Request, *GetProduct.Response](GetProduct.NewHandler(productRepository))
	_ = mediatr.RegisterRequestHandler[*GetProducts.Request, *GetProducts.Response](GetProducts.NewHandler(productRepository))
	_ = mediatr.RegisterRequestHandler[*AddProduct.Request, *AddProduct.Response](AddProduct.NewHandler(productRepository))
	_ = mediatr.RegisterRequestHandler[*DeleteProduct.Request, *DeleteProduct.Response](DeleteProduct.NewHandler(productRepository))
	_ = mediatr.RegisterRequestHandler[*PatchProduct.Request, *PatchProduct.Response](PatchProduct.NewHandler(productRepository))

	e.GET("/products", func(c echo.Context) error {
		request := GetProducts.NewRequest()
		response, err := mediatr.Send[*GetProducts.Request, *GetProducts.Response](c.Request().Context(), request)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, response)
	})

	e.GET("/products/:id", func(c echo.Context) error {
		productId := c.Param("id")
		productIdIntegerized, _ := strconv.Atoi(productId)
		productIdUintegerized := uint(productIdIntegerized) //TODO: this parsing should be done somehow better...
		request := GetProduct.NewRequest(productIdUintegerized)

		response, err := mediatr.Send[*GetProduct.Request, *GetProduct.Response](c.Request().Context(), request)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, response)
	})

	e.POST("/products", func(c echo.Context) error {
		//TODO: use bind with some DTO
		newProduct := Models.NewProductBinding{}
		err := c.Bind(&newProduct)
		if err != nil {
			return c.String(http.StatusBadRequest, "Error during binding.")
		}

		request := AddProduct.NewRequest(newProduct.Id, newProduct.Name)

		response, err := mediatr.Send[*AddProduct.Request, *AddProduct.Response](c.Request().Context(), request)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, response)
	})

	e.DELETE("/products/:id", func(c echo.Context) error {
		productId := c.Param("id")
		productIdIntegerized, _ := strconv.Atoi(productId)
		productIdUintegerized := uint(productIdIntegerized) //TODO: this parsing should be done somehow better...
		request := DeleteProduct.NewRequest(productIdUintegerized)

		response, err := mediatr.Send[*DeleteProduct.Request, *DeleteProduct.Response](c.Request().Context(), request)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, response)
	})

	e.PATCH("/products/:id", func(c echo.Context) error {
		productId := c.Param("id")
		productIdIntegerized, _ := strconv.Atoi(productId)
		productIdUintegerized := uint(productIdIntegerized) //TODO: this parsing should be done somehow better...

		patchProductBinding := Models.PatchProductBinding{}
		err := c.Bind(&patchProductBinding)
		if err != nil {
			return c.String(http.StatusBadRequest, "Error during binding.")
		}
		request := PatchProduct.NewRequest(productIdUintegerized, patchProductBinding)
		response, err := mediatr.Send[*PatchProduct.Request, *PatchProduct.Response](c.Request().Context(), request)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, response)
	})
}
