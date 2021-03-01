package presenter

import (
	"github.com/Topi99/academy-go-q12021/domain/model"
	"github.com/Topi99/academy-go-q12021/usecase/presenter"
)

type pokemonPresenter struct{}

// NewPokemonPresenter creates new PokemonPresenter
func NewPokemonPresenter() presenter.PokemonPresenter {
	return &pokemonPresenter{}
}

func (pp *pokemonPresenter) ResponsePokemon(p *model.Pokemon) *model.Pokemon {
	return p
}
