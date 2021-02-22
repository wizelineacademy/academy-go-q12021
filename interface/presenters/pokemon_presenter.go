package presenter

import "pokeapi/domain/model"

type pokemonPresenter struct {
}

type PokemonPresenter interface {
	ResponsePokemons(us []*model.Pokemon) []*model.Pokemon
}

func NewPokemonPresenter() PokemonPresenter {
	return &pokemonPresenter{}
}

func (pp *pokemonPresenter) ResponsePokemons(ps []*model.Pokemon) []*model.Pokemon {
	for _, p := range ps {
		p.Name = "Pokemon " + p.Name
	}
	return ps
}
