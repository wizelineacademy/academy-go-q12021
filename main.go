package main

import (
	"log"
	"net/http"

	"github.com/grethelBello/academy-go-q12021/config"
	"github.com/grethelBello/academy-go-q12021/constant"
	"github.com/grethelBello/academy-go-q12021/controller"
)

func main() {
	pokemonController, error := controller.NewPokemonController()
	if error != nil {
		log.Fatal(error)
	}
	http.HandleFunc("/first-delivery/pokemons", pokemonController.GetCsvPokemons)
	http.HandleFunc("/second-delivery/pokemons", pokemonController.GetDynamicPokemons)
	http.HandleFunc("/final-delivery/pokemons", pokemonController.GetCurrentPokemons)

	serverPort, serverError := config.GetEnvVar(constant.ServerPortVarName)
	if serverError != nil {
		log.Fatal(serverError)
	}
	initError := http.ListenAndServe(":"+serverPort, nil)
	log.Println(initError)

}
