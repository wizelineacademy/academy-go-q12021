package router

import (
	"net/http"
	"pokeapi/controller"

	"github.com/gorilla/mux"
)

type Router struct {
	controller controller.NewPokemonController
}

type IRouter interface {
	InitRouter() *mux.Router
}

func New(c controller.NewPokemonController) *Router {
	return &Router{c}
}

func (router *Router) InitRouter() *mux.Router {
	r := mux.NewRouter()

	r.Methods(http.MethodGet).
		Path("/pokemons/concurrency/{type}").
		Queries("items_per_worker", "{[0-9]+}").
		Queries("items", "{[0-9]+}").
		HandlerFunc(router.controller.GetPokemonConcurrently)
	r.Methods(http.MethodGet).
		Path("/pokemons/concurrency/{type}").
		Queries("items", "{[0-9]+}").
		HandlerFunc(router.controller.GetPokemonConcurrently)
	r.HandleFunc("/pokemons/external", router.controller.GetPokemonsFromExternalAPI).Methods("GET")
	r.HandleFunc("/pokemons/{id}", router.controller.GetPokemon).Methods("GET")
	r.HandleFunc("/pokemons", router.controller.GetPokemons).Methods("GET")
	r.HandleFunc("/", router.controller.Index)

	return r
}
