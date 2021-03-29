package main

import (
	"log"
	"net/http"

	"pokeapi/controller"
	"pokeapi/router"
	csvservice "pokeapi/service/csv"
	httpservice "pokeapi/service/http"
	"pokeapi/usecase"
)

func main() {

	csvService := csvservice.New()
	httpService := httpservice.New()
	usecase := usecase.New(csvService, httpService)
	controller := controller.New(usecase)

	router := router.New(controller)
	r := router.InitRouter()
	log.Fatal(http.ListenAndServe(":3000", r))
}
