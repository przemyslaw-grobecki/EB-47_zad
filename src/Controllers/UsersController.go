package Controllers

import (
	"4_zad/Bussiness/Commands/LoginUser"
	"4_zad/Bussiness/Commands/RegisterUser"
	"4_zad/Models"
	"4_zad/Services"
	"github.com/labstack/echo/v4"
	"github.com/mehdihadeli/go-mediatr"
	"net/http"
)

func BootstrapUsersController(e *echo.Echo, userRepository Services.IRepository[Models.User]) {
	_ = mediatr.RegisterRequestHandler[*RegisterUser.Request, *RegisterUser.Response](RegisterUser.NewHandler(userRepository))
	_ = mediatr.RegisterRequestHandler[*LoginUser.Request, *LoginUser.Response](LoginUser.NewHandler(userRepository))

	e.POST("/users/register", func(c echo.Context) error {
		newUser := Models.UserBinding{}
		err := c.Bind(&newUser)
		if err != nil {
			return c.String(http.StatusBadRequest, "Error during binding.")
		}

		request := RegisterUser.NewRequest(newUser.Email, newUser.Password)

		response, err := mediatr.Send[*RegisterUser.Request, *RegisterUser.Response](c.Request().Context(), request)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, response)
	})

	e.POST("/users/login", func(c echo.Context) error {
		user := Models.UserBinding{}
		err := c.Bind(&user)
		if err != nil {
			return c.String(http.StatusBadRequest, "Error during binding.")
		}

		request := LoginUser.NewRequest(user.Email, user.Password)
		response, err := mediatr.Send[*LoginUser.Request, *LoginUser.Response](c.Request().Context(), request)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, response)
	})
}
