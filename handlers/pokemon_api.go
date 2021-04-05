package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// GetPokemonsFromAPI is the handler to get all the pokemons by quantity
func (p *Pokemons) GetPokemonsFromAPI(w http.ResponseWriter, r *http.Request) {
	logger := p.logger.WithFields(logrus.Fields{
		"func": "Get Pokemons",
	})
	logger.Debug("in")
	quantity := r.FormValue("quantity")
	if quantity == "" {
		quantity = "1"
	}
	pokemons, err := p.useCase.GetPokemonsFromAPI(quantity)
	if err != nil {
		p.render.JSON(w, http.StatusInternalServerError, pokemons)
	}

	p.render.JSON(w, http.StatusOK, pokemons)

}

// GetPokemonsFromAPI is the handler to get the pokemon by id
func (p *Pokemons) GetPokemonFromAPI(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pokemonID := vars["id"]
	logger := p.logger.WithFields(logrus.Fields{
		"func": "Get Pokemon",
		"id":   pokemonID,
	})
	logger.Debug("in")

	pokemon, err := p.useCase.GetPokemonFromAPI(pokemonID)

	if err == nil {
		p.render.JSON(w, http.StatusOK, pokemon)
	}
	p.render.JSON(w, http.StatusNotFound, pokemon)
}
