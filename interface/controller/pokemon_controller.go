package controller

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"

	"github.com/AlejandroSeguraWIZ/academy-go-q12021/usecase/interactor"
)

type pokemonController struct {
	pokemonInteractor interactor.PokemonInteractor
}

type PokemonController interface {
	GetPokemons(echo.Context) error
}

func NewPokemonController(pi interactor.PokemonInteractor) PokemonController {
	return &pokemonController{
		pokemonInteractor: pi,
	}
}

func (pc *pokemonController) GetPokemons(ctx echo.Context) error {
	resp, err := pc.pokemonInteractor.Get()
	if err != nil {
		fmt.Print(err)
		ctx.JSON(http.StatusInternalServerError, nil)
	}
	return ctx.JSON(http.StatusOK, resp)
}
