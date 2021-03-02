package router

import (
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
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/", router.controller.Index)
	r.HandleFunc("/pokemons", router.controller.GetPokemons).Methods("GET")
	r.HandleFunc("/pokemons/{id}", router.controller.GetPokemon).Methods("GET")
	r.HandleFunc("/pokemons/external", router.controller.GetPokemonsFromExternalAPI).Methods("GET")

	return r
}
