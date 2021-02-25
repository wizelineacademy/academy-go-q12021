package csv

import (
	"bootcamp/utils"
	"bootcamp/domain/model"
)

func GetPokemon() (model.PokemonList, error) {
	pokeList, err := utils.ReadCSV()
	return pokeList, err
}