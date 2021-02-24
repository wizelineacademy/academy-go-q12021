package main

import (
	"log"
	"net/http"
	"pokeapi/controller"
	"pokeapi/router"
	"pokeapi/service"
	"pokeapi/usecase"
)

func main() {

	httpService := service.NewHttp()
	service := service.NewCsv()
	usecase := usecase.New(service, httpService)
	controller := controller.New(usecase)

	router := router.New(controller)
	r := router.InitRouter()
	log.Fatal(http.ListenAndServe(":3000", r))
}
