package presenter

import "pokeapi/domain/model"

type PokemonPresenter interface {
	ResponsePokemons(p []*model.Pokemon) []*model.Pokemon
}
