package router

import (
  // "encoding/json"
  "fmt"
  "github.com/labstack/echo/v4"
  "github.com/labstack/echo/v4/middleware"


	"net/http"
)

func homePage(c echo.Context) error {
  fmt.Println("endpont hit: AllArticles!")

  return c.String(http.StatusOK, "HOLA")
}

func NewRouter(e *echo.Echo) *echo.Echo {
  e.Use(middleware.Logger())
  e.Use(middleware.Recover())

  e.GET("/", homePage)

  return e
}
