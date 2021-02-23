package routes

import (
	"modules"
	"github.com/gorilla/mux"
	"net/http"
)

type Route struct{
	Name string
	Method string
	Pattern string
	HandleFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Name(route.Name).
			Methods(route.Method).
			Path(route.Pattern).
			Handler(route.HandleFunc)
	}
	return router
}

var get = "GET"
var post = "POST"
var pokemonPath = "/pokemon"
var csv = "/csv"

var routes = Routes{
	Route{
		"HelloWorld",
		get,
		"/",
		modules.HelloWorld,
	},
	Route{
		"GetPokemonListCsv",
		get,
		csv + pokemonPath,
		modules.GetPokemonListCsv,
	},
	Route{
		"GetPokemonCsv",
		get,
		csv + pokemonPath + "/{id}",
		modules.GetPokemonCsv,
	},
	Route{
		"AddPokemon",
		post,
		pokemonPath,
		modules.AddPokemon,
	},
}