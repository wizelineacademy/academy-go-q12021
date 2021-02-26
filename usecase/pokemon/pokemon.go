package pokemon

import (
	"bootcamp/utils"
	"bootcamp/domain/model"
	"bootcamp/service/db"
	"io"
)

func GetPokemon() (model.PokemonList, error) {
	pokemonList, err := db.GetPokemon()
	return pokemonList, err
}

func GetPokemonById(params map[string]string) (model.Pokemon, error) {
	var pokemon model.Pokemon
	objectId, err := utils.GetObjectIdFromParams(params)

	if err == nil {
		pokemon, err = db.GetPokemonById(objectId)
	}

	return pokemon, err
}

func AddPokemon(reader io.ReadCloser) (model.Pokemon, error) {
	pokemon, err := utils.GetPokemonFromReader(reader)
	
	if err == nil {
		pokemon, err = db.AddPokemon(pokemon)
	}

	return pokemon, err
}

func UpdatePokemon(params map[string]string, reader io.ReadCloser) (model.Pokemon, error) {
	var pokemon model.Pokemon
	pokemon, err := utils.GetPokemonFromReader(reader)
	objectId, err := utils.GetObjectIdFromParams(params)

	if err == nil {
		pokemon, err = db.UpdatePokemon(objectId, pokemon)
	}

	return pokemon, err
}

func DeletePokemon(params map[string]string) (model.Pokemon, error) {
	var pokemon model.Pokemon
	objectId, err := utils.GetObjectIdFromParams(params)

	if err == nil {
		pokemon, err = db.DeletePokemon(objectId)
	}

	return pokemon, err
}
