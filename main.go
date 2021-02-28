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

	csvService := service.New()
	usecase := usecase.New(csvService)
	controller := controller.New(usecase)

	router := router.New(controller)
	r := router.InitRouter()
	log.Fatal(http.ListenAndServe(":3000", r))
}
