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
	router.HandleFunc("/pokemon/workers", c.GetPokemonsWithWorkers).
		Queries("type", "{type:[a-z]+}", "items", "{items:[0-9]+}",
			"items_per_workers","{items_per_workers:[0-9]+}")
	return router
}
