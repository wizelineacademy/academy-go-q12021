package main

import (
	"log"
	"net/http"
	"pokeapi/controllers"
	"pokeapi/services"
	"pokeapi/usecases"
)

func main() {

	service := services.NewService()
	httpService := services.NewHttpService()
	usecase := usecases.NewUseCase(service, httpService)
	controller := controllers.NewPokemonController(usecase)

	router := NewRouter(controller)
	r := router.InitRouter()
	log.Fatal(http.ListenAndServe(":3000", r))
}
