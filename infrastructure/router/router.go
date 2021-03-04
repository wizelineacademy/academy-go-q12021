package router

import (
	"bootcamp/interface/controller"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func NewRouter(e *echo.Echo, c controller.AppController) *echo.Echo {

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/items", func(context echo.Context) error { return c.Item.GetItems(context) })
	e.GET("/items/:id", func(context echo.Context) error { return c.Item.GetItem(context) })

	return e
}