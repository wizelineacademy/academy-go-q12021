package main

import (
	"github.com/oscarSantoyo/academy-go-q12021/container"
	"github.com/oscarSantoyo/academy-go-q12021/router"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
)

func init() {
	container.Connect()
}

func main() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		message := "There was an error reading the configuration file"
		log.Fatal(message, err)
		panic(message)
	}
	startServer()
}

func startServer() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	router.SetRouter(e)
}
