package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/wizelineacademy/academy-go/config"
	"github.com/wizelineacademy/academy-go/constant"
	"github.com/wizelineacademy/academy-go/controller"
)

func main() {
	pokemonController, error := controller.NewPokemonController()
	if error != nil {
		panic(error)
	}
	http.HandleFunc("/pokemons", pokemonController.GetPokemons)

	serverPort, serverError := config.GetEnvVar(constant.ServerPortVarName)
	if serverError != nil {
		log.Fatal(serverError)
	}
	initError := http.ListenAndServe(":"+serverPort, nil)
	fmt.Println(initError)

}
