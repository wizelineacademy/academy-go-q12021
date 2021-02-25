package router

import (
	"bootcamp/domain/model"
	"bootcamp/controller/hello"
	"bootcamp/controller/csv"
	"bootcamp/controller/pokemon"
	// "modules"
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
	// model.Route{
	// 	"AddPokemon",
	// 	http.MethodPost,
	// 	pokemonPath,
	// 	modules.AddPokemon,
	// },
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
	// model.Route{
	// 	"UpdatePokemon",
	// 	http.MethodPut,
	// 	pokemonPathWithId,
	// 	modules.UpdatePokemon,
	// },
	// model.Route{
	// 	"DeletePokemon",
	// 	http.MethodDelete,
	// 	pokemonPathWithId,
	// 	modules.DeletePokemon,
	// },
}