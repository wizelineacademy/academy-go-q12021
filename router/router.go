package router

import (
	"model"
	"bootcamp/controller/hello"
	"modules"
	"github.com/gorilla/mux"
	"net/http"
)

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

var pokemonPath = "/pokemon"
var id = "{id}"
var pokemonPathWithId = pokemonPath + id
var csvPokemonPath = pokemonPath + "/csv"

var routes = model.Routes{
	model.Route{
		"HelloWorld",
		http.MethodGet,
		"/",
		hello.HelloWorld,
	},
	model.Route{
		"GetPokemonListCsv",
		http.MethodGet,
		csvPokemonPath,
		modules.GetPokemonListCsv,
	},
	model.Route{
		"GetPokemonCsv",
		http.MethodGet,
		csvPokemonPath + id,
		modules.GetPokemonCsv,
	},
	model.Route{
		"AddPokemon",
		http.MethodPost,
		pokemonPath,
		modules.AddPokemon,
	},
	model.Route{
		"GetPokemonList",
		http.MethodGet,
		pokemonPath,
		modules.GetPokemonList,
	},
	model.Route{
		"GetPokemon",
		http.MethodGet,
		pokemonPathWithId,
		modules.GetPokemon,
	},
	model.Route{
		"UpdatePokemon",
		http.MethodPut,
		pokemonPathWithId,
		modules.UpdatePokemon,
	},
	model.Route{
		"DeletePokemon",
		http.MethodDelete,
		pokemonPathWithId,
		modules.DeletePokemon,
	},
}