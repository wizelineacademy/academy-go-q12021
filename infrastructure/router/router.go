package router

import (
	"fmt"

	controller "github.com/ToteEmmanuel/academy-go-q12021/controller"
	"github.com/gorilla/mux"
)

func NewRouter(c controller.AppController) *mux.Router {
	fmt.Printf("Powering up...\n")
	router := mux.NewRouter()
	router.HandleFunc("/pokemon", c.GetPokemons)
	router.HandleFunc("/pokemon/{id:[0-9]+}", c.GetPokemon)
	router.HandleFunc("/pokemon/{id:[0-9]+}/catch", c.CatchPokemon)
	return router
}
