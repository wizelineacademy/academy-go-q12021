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
	http.HandleFunc("/pokemons", pokemonController.GetPokemons)

	serverPort, serverError := config.GetEnvVar(constant.ServerPortVarName)
	if serverError != nil {
		log.Fatal(serverError)
	}
	initError := http.ListenAndServe(":"+serverPort, nil)
	log.Println(initError)

}
