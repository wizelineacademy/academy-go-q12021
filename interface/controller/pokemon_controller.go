package controller

import (
	"net/http"

	"github.com/Topi99/academy-go-q12021/usecase/interactor"
)

type pokemonController struct {
	pokemonInteractor interactor.PokemonInteractor
}

// PokemonController interface
type PokemonController interface {
	GetOne(c Context, id uint) error
}

// NewPokemonController returns new pokemonController
func NewPokemonController(p interactor.PokemonInteractor) PokemonController {
	return &pokemonController{p}
}

func (po *pokemonController) GetOne(c Context, id uint) error {
	p, err := po.pokemonInteractor.GetOne(id)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, p)
}
