package main

import (
	"pokeapi/controllers"

	"github.com/gorilla/mux"
)

type Router struct {
	controller controllers.IPokemonController
}

type IRouter interface {
	InitRouter() *mux.Router
}

func NewRouter(c controllers.IPokemonController) *Router {
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
