package main

import (
	"net/http"

	"./config"
	"./controller"
)

var serverPort = config.ServerPort

func main() {
	pokemonRoutes := controller.GetPokemonRoutes()
	for path, handler := range pokemonRoutes {
		http.HandleFunc(path, handler)
	}
	http.ListenAndServe(":8080", nil)
}
