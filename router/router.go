package router

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/AlejandroSeguraWIZ/academy-go-q12021/interface/controller"
)

func NewRouter(e *echo.Echo, c controller.AppController) *echo.Echo {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/pokemons", func(ctx echo.Context) error {
		return c.GetPokemons(ctx)
	})
	return e
}
