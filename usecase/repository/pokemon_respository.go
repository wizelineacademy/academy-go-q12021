package repository

import "pokeapi/domain/model"

type PokemonRepository interface {
	FindAll(p []*model.Pokemon) ([]*model.Pokemon, error)
}
