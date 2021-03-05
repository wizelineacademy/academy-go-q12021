package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/ToteEmmanuel/academy-go-q12021/usecase/interactor"
	"github.com/gorilla/mux"
)

type pokeController struct {
	pokeInteractor interactor.PokeInteractor
}

type PokeController interface {
	GetPokemon(http.ResponseWriter, *http.Request)
	GetPokemons(http.ResponseWriter, *http.Request)
}

func NewPokeController(pokeInteractor interactor.PokeInteractor) PokeController {
	return &pokeController{pokeInteractor}
}

func (pI *pokeController) GetPokemon(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	p, err := pI.pokeInteractor.Get(int32(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(p)
}

func (pI *pokeController) GetPokemons(w http.ResponseWriter, r *http.Request) {
	p := pI.pokeInteractor.GetAll()
	json.NewEncoder(w).Encode(p)
}
