package router

import (
	"bootcamp/infrastructure/validator"
	"bootcamp/interface/controller"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// NewRouter sets up the Echo routing
func NewRouter(e *echo.Echo, c controller.AppController) *echo.Echo {

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Validator = validator.NewCustomValidator()

	e.GET("/items", func(context echo.Context) error { return c.Item.GetItems(context) })
	e.GET("/items/:id", func(context echo.Context) error { return c.Item.GetItem(context) })
	e.GET("/jokes", func(context echo.Context) error { return c.Joke.GetJokes(context) })

	return e
}