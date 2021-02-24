package router

import (
	"pokeapi/controller"

	"github.com/gorilla/mux"
)

type Router struct {
	controller controller.IPokemonController
}

type IRouter interface {
	InitRouter() *mux.Router
}

func New(c controller.IPokemonController) *Router {
	return &Router{c}
}

func (router *Router) InitRouter() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/", router.controller.Index)
	r.HandleFunc("/pokemons", router.controller.GetPokemons).Methods("GET")
	r.HandleFunc("/pokemons/{id}", router.controller.GetPokemon).Methods("GET")
	r.HandleFunc("/external/pokemons", router.controller.GetPokemonsFromExternalAPI).Methods("GET")

	return r
}
