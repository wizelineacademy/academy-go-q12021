package presenter

import "github.com/ToteEmmanuel/academy-go-q12021/domain/model"

type PokePresenter interface {
	ResponsePokemon(*model.Pokemon) *model.Pokemon
	ResponsePokemons([]*model.Pokemon) []*model.Pokemon
}
