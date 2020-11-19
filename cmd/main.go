package main

import (
	"github.com/AlonSerrano/GolangBootcamp/pkg/handlers"
	"github.com/labstack/echo/v4"
	"log"
)

func main() {
	e := echo.New()
	v1 := e.Group("/api/v1")
	h := handlers.NewPerissonConnectorHandler()
	direcciones := v1.Group("/direcciones")
	{
		direcciones.GET("/populate", h.HandlePopulateZipCodes)
		direcciones.GET("/search/:zipCode", h.HandleSearchZipCodes)
	}

	serverAddress := "localhost:8080"
	log.Printf("server started at %s\n", serverAddress)
	e.Logger.Fatal(e.Start(serverAddress))
}
