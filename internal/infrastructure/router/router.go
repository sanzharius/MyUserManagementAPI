package router

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"myAPIProject/internal/adapter/controller"
)

func NewRouter(e *echo.Echo, c controller.AppController) *echo.Echo {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Validator = &controller.CustomValidator{Validator: validator.New()}

	e.GET("/users", func(context echo.Context) error { return c.UserController.GetUsers(context) })
	e.GET("/user/:id", func(context echo.Context) error { return c.UserController.GetUser(context) })
	e.POST("/user", func(context echo.Context) error { return c.UserController.CreateUser(context) } /*, c.UserController.BasicAuth()*/)
	e.PUT("/user/:id", func(context echo.Context) error { return c.UserController.UpdateUser(context) }, c.UserController.BasicAuth())
	e.DELETE("/user/:id", func(context echo.Context) error { return c.UserController.DeleteUser(context) }, c.UserController.BasicAuth())

	return e
}
