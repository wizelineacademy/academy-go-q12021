package main

import (
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

var routes = Routes{
	Route{
		"HelloWorld",
		"GET",
		"/",
		HelloWorld,
	},
	Route{
		"GetPokemonList",
		"GET",
		"/pokemon",
		GetPokemonList,
	},
	Route{
		"GetPokemon",
		"GET",
		"/pokemon/{id}",
		GetPokemon,
	},
	Route{
		"AddPokemon",
		"POST",
		"/pokemon",
		AddPokemon,
	},
}