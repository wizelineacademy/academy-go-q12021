package router

import (
	"net/http"
	"bootcamp/controller/csv"
	"bootcamp/controller/hello"
	"bootcamp/controller/pokemon"
	"bootcamp/domain/model"
	"github.com/gorilla/mux"
)

/*
NewRouter implements a gorilla/mux router with the routes of the API
*/
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
var id = "/{id}"
var pokemonPathWithId = pokemonPath + id
var csvPokemonPath = "/csv" + pokemonPath

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
		csv.GetPokemon,
	},
	model.Route{
		"GetPokemonCsv",
		http.MethodGet,
		csvPokemonPath + id,
		csv.GetPokemon,
	},
	model.Route{
		"AddPokemon",
		http.MethodPost,
		pokemonPath,
		pokemon.AddPokemon,
	},
	model.Route{
		"GetPokemon",
		http.MethodGet,
		pokemonPath,
		pokemon.GetPokemon,
	},
	model.Route{
		"GetPokemonById",
		http.MethodGet,
		pokemonPathWithId,
		pokemon.GetPokemon,
	},
	model.Route{
		"UpdatePokemon",
		http.MethodPut,
		pokemonPathWithId,
		pokemon.UpdatePokemon,
	},
	model.Route{
		"DeletePokemon",
		http.MethodDelete,
		pokemonPathWithId,
		pokemon.DeletePokemon,
	},
}