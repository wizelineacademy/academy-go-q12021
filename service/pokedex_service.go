package service

import (
    "pokedex/model"
)

type PokedexService interface {
    Init() error
    GetPokemonById(id int) (model.Pokemon, error)
}
