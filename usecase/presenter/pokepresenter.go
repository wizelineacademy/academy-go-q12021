package presenter

import "github.com/ToteEmmanuel/academy-go-q12021/domain/model"

type pokePresenter struct {
}

type PokePresenter interface {
	ResponsePokemon(*model.Pokemon) *model.Pokemon
	ResponsePokemons([]*model.Pokemon) []*model.Pokemon
}

func NewPokePresenter() PokePresenter {
	return &pokePresenter{}
}

func (pP *pokePresenter) ResponsePokemon(pokething *model.Pokemon) *model.Pokemon {
	return pokething
}

func (pP *pokePresenter) ResponsePokemons(pokethings []*model.Pokemon) []*model.Pokemon {
	return pokethings
}
