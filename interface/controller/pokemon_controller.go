package controller

import (
	"net/http"
	"strconv"

	"github.com/Topi99/academy-go-q12021/usecase/interactor"
)

type pokemonController struct {
	pokemonInteractor interactor.PokemonInteractor
}

// PokemonController interface
type PokemonController interface {
	GetOne(c Context) error
}

// NewPokemonController returns new pokemonController
func NewPokemonController(p interactor.PokemonInteractor) PokemonController {
	return &pokemonController{p}
}

func (po *pokemonController) GetOne(c Context) error {
	idU64, err := strconv.ParseUint(c.Param("id"), 10, 32)

	if err != nil {
		return err
	}

	id := uint(idU64)

	p, err := po.pokemonInteractor.GetOne(id)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, p)
}
