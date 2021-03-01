package router

import (
	"github.com/Topi99/academy-go-q12021/interface/controller"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// NewRouter creates new echo.Echo router
func NewRouter(e *echo.Echo, c controller.AppController) *echo.Echo {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/pokemon/:id", func(context echo.Context) error { return c.Pokemon.GetOne(context) })

	return e
}
