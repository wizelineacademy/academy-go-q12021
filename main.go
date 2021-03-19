package main

import (
	"net/http"

	"github.com/wizelineacademy/academy-go/config"
	"github.com/wizelineacademy/academy-go/controller"
	"github.com/wizelineacademy/academy-go/data"
	"github.com/wizelineacademy/academy-go/model"
	"github.com/wizelineacademy/academy-go/service"
)

var serverPort = config.GetServerPort()
var routes map[string]func(w http.ResponseWriter, r *http.Request)

func init() {
	pokemonSource := config.GetPokemonSource()
	pokemonDataSource := data.CsvSource(pokemonSource)
	pokemonDataService := service.PokemonDataService(make(map[int]model.Pokemon))
	pokemonDataService.Init(pokemonDataSource)
	routes = controller.GetPokemonRoutes(pokemonDataService)
}

func main() {

	for path, handler := range routes {
		http.HandleFunc(path, handler)
	}
	http.ListenAndServe(":"+serverPort, nil)
}
