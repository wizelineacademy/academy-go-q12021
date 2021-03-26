package main

import (
	"fmt"
	"net/http"

	"github.com/wizelineacademy/academy-go/config"
	"github.com/wizelineacademy/academy-go/constant"
	"github.com/wizelineacademy/academy-go/controller"
)

var serverPort = config.GetEnvVar(constant.ServerPortVarName)

func main() {
	pokemonCsvPath := config.GetEnvVar(constant.PokemonSourceVarName)
	pokemonApi := config.GetEnvVar(constant.PokemonServiceVarName)
	pokemonController, error := controller.NewPokemonController(pokemonCsvPath, pokemonApi)
	if error != nil {
		panic(error)
	}

	http.HandleFunc("/pokemons", pokemonController.GetPokemons)
	serverError := http.ListenAndServe(":"+serverPort, nil)
	fmt.Println(serverError)

}
