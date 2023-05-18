package Controllers

import (
	"4_zad/Bussiness/Commands/AddBasket"
	"4_zad/Bussiness/Commands/AddOrder"
	"4_zad/Bussiness/Commands/BuyBasket"
	"4_zad/Bussiness/Commands/DeleteBasket"
	"4_zad/Bussiness/Commands/PatchBasket"
	"4_zad/Bussiness/Commands/PatchOrder"
	"4_zad/Bussiness/Queries/GetBasket"
	"4_zad/Bussiness/Queries/GetBaskets"
	"4_zad/Models"
	"4_zad/Services"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/mehdihadeli/go-mediatr"
	"io/ioutil"
	"net/http"
	"strconv"
)

func BootstrapBasketsController(e *echo.Echo, basketRepository Services.IRepository[Models.Basket], productRepository Services.IRepository[Models.Product]) {
	_ = mediatr.RegisterRequestHandler[*GetBasket.Request, *GetBasket.Response](GetBasket.NewHandler(basketRepository))
	_ = mediatr.RegisterRequestHandler[*GetBaskets.Request, *GetBaskets.Response](GetBaskets.NewHandler(basketRepository))
	_ = mediatr.RegisterRequestHandler[*AddBasket.Request, *AddBasket.Response](AddBasket.NewHandler(basketRepository))
	_ = mediatr.RegisterRequestHandler[*DeleteBasket.Request, *DeleteBasket.Response](DeleteBasket.NewHandler(basketRepository))
	_ = mediatr.RegisterRequestHandler[*PatchBasket.Request, *PatchBasket.Response](PatchBasket.NewHandler(basketRepository))
	_ = mediatr.RegisterRequestHandler[*AddOrder.Request, *AddOrder.Response](AddOrder.NewHandler(basketRepository, productRepository))
	_ = mediatr.RegisterRequestHandler[*BuyBasket.Request, *BuyBasket.Response](BuyBasket.NewHandler(basketRepository, productRepository))

	e.GET("/baskets", func(c echo.Context) error {
		request := GetBaskets.NewRequest()
		response, err := mediatr.Send[*GetBaskets.Request, *GetBaskets.Response](c.Request().Context(), request)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, response)
	})

	e.GET("/baskets/:id", func(c echo.Context) error {
		basketId := c.Param("id")
		basketIdIntegerized, _ := strconv.Atoi(basketId)
		basketIdUintegerized := uint(basketIdIntegerized) //TODO: this parsing should be done somehow better...
		request := GetBasket.NewRequest(basketIdUintegerized)

		response, err := mediatr.Send[*GetBasket.Request, *GetBasket.Response](c.Request().Context(), request)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, response)
	})

	e.POST("/baskets", func(c echo.Context) error {
		newBasket := Models.NewBasketBinding{}
		err := c.Bind(&newBasket)
		if err != nil {
			return c.String(http.StatusBadRequest, "Error during binding.")
		}

		request := AddBasket.NewRequest(newBasket.Id)

		response, err := mediatr.Send[*AddBasket.Request, *AddBasket.Response](c.Request().Context(), request)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, response)
	})

	e.POST("/baskets/:id/orders", func(c echo.Context) error {
		basketId := c.Param("id")
		basketIdIntegerized, _ := strconv.Atoi(basketId)
		basketIdUintegerized := uint(basketIdIntegerized) //TODO: this parsing should be done somehow better...
		newOrder := Models.NewOrderBinding{}
		err := c.Bind(&newOrder)
		if err != nil {
			return c.String(http.StatusBadRequest, "Error during binding.")
		}

		request := AddOrder.NewRequest(basketIdUintegerized, newOrder.ProductId, newOrder.Quantity)

		response, err := mediatr.Send[*AddOrder.Request, *AddOrder.Response](c.Request().Context(), request)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, response)
	})

	e.POST("/baskets/:id/buy", func(c echo.Context) error {
		basketId := c.Param("id")
		basketIdIntegerized, _ := strconv.Atoi(basketId)
		basketIdUintegerized := uint(basketIdIntegerized) //TODO: this parsing should be done somehow better...

		buyBasket := Models.BuyBasketBinding{}

		body, err := ioutil.ReadAll(c.Request().Body)
		if err != nil {
			return err
		}

		err = json.Unmarshal(body, &buyBasket)
		if err != nil {
			return err
		}

		request := BuyBasket.NewRequest(basketIdUintegerized, buyBasket.Orders)
		response, err := mediatr.Send[*BuyBasket.Request, *BuyBasket.Response](c.Request().Context(), request)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, response)
	})

	e.DELETE("/baskets/:id", func(c echo.Context) error {
		basketId := c.Param("id")
		basketIdIntegerized, _ := strconv.Atoi(basketId)
		basketIdUintegerized := uint(basketIdIntegerized) //TODO: this parsing should be done somehow better...
		request := DeleteBasket.NewRequest(basketIdUintegerized)

		response, err := mediatr.Send[*DeleteBasket.Request, *DeleteBasket.Response](c.Request().Context(), request)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, response)
	})

	e.PATCH("/baskets/:id", func(c echo.Context) error {
		basketId := c.Param("id")
		basketIdIntegerized, _ := strconv.Atoi(basketId)
		basketIdUintegerized := uint(basketIdIntegerized) //TODO: this parsing should be done somehow better...

		patchBasketBinding := Models.PatchBasketBinding{}
		err := c.Bind(&patchBasketBinding)
		if err != nil {
			return c.String(http.StatusBadRequest, "Error during binding.")
		}
		request := PatchBasket.NewRequest(basketIdUintegerized, patchBasketBinding)
		response, err := mediatr.Send[*PatchBasket.Request, *PatchBasket.Response](c.Request().Context(), request)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, response)
	})

	e.PATCH("/baskets/:basketId/orders/:orderId", func(c echo.Context) error {
		basketId := c.Param("basketId")
		basketIdIntegerized, _ := strconv.Atoi(basketId)
		basketIdUintegerized := uint(basketIdIntegerized) //TODO: this parsing should be done somehow better...

		orderId := c.Param("orderId")
		orderIdIntegerized, _ := strconv.Atoi(orderId)
		orderIdUintegerized := uint(orderIdIntegerized)

		patchOrderBinding := Models.PatchOrderBinding{}
		err := c.Bind(&patchOrderBinding)
		if err != nil {
			return c.String(http.StatusBadRequest, "Error during binding.")
		}
		request := PatchOrder.NewRequest(basketIdUintegerized, orderIdUintegerized, patchOrderBinding)
		response, err := mediatr.Send[*PatchOrder.Request, *PatchOrder.Response](c.Request().Context(), request)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, response)
	})
}
