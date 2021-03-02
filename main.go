package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/oscarSantoyo/academy-go-q12021/container"
	"github.com/oscarSantoyo/academy-go-q12021/router"
)

func init() {
	container.Connect()
}

func main() {
	startServer()
}

func startServer() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	router.SetRouter(e)

	e.Logger.Fatal(e.Start(":8080"))
}
