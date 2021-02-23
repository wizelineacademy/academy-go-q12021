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

var routes = Routes{
	Route{
		"HelloWorld",
		get,
		"/",
		modules.HelloWorld,
	},
	Route{
		"GetPokemonList",
		get,
		"/pokemon",
		modules.GetPokemonList,
	},
	Route{
		"GetPokemon",
		get,
		"/pokemon/{id}",
		modules.GetPokemon,
	},
	Route{
		"AddPokemon",
		post,
		"/pokemon",
		modules.AddPokemon,
	},
}