package controller

import (
	"net/http"
	"pokeapi/domain/model"
	"pokeapi/usecase/interactor"
)

type pokemonController struct {
	pokemonInteractor interactor.PokemonInteractor
}

type PokemonController interface {
	GetPokemons(c Context) error
}

func NewPokemonController(ps interactor.PokemonInteractor) PokemonController {
	return &pokemonController{ps}
}

func (pc *pokemonController) GetPokemons(c Context) error {
	var p []*model.Pokemon

	p, err := pc.pokemonInteractor.Get(p)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, p)
}
