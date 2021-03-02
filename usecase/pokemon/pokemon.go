package pokemon

import (
	"io"	
	"bootcamp/domain/model"
	"bootcamp/service/db"
	"bootcamp/utils"
)

/*
GetPokemon retrieve all existent Pokemon from the database
*/
func GetPokemon() (model.PokemonList, error) {
	pokemonList, err := db.GetPokemon()
	return pokemonList, err
}

/*
GetPokemonById retrieves Pokemon information that matches with a given id from the database
*/
func GetPokemonById(params map[string]string) (model.Pokemon, error) {
	var pokemon model.Pokemon
	objectId, err := utils.GetObjectIdFromParams(params)

	if err == nil {
		pokemon, err = db.GetPokemonById(objectId)
	}

	return pokemon, err
}

/*
AddPokemon inserts Pokemon information in the database
*/
func AddPokemon(reader io.ReadCloser) (model.Pokemon, error) {
	pokemon, err := utils.GetPokemonFromReader(reader)
	
	if err == nil {
		pokemon, err = db.AddPokemon(pokemon)
	}

	return pokemon, err
}

/*
UpdatePokemon updates Pokemon information in the database
*/
func UpdatePokemon(params map[string]string, reader io.ReadCloser) (model.Pokemon, error) {
	var pokemon model.Pokemon
	pokemon, err := utils.GetPokemonFromReader(reader)
	objectId, err := utils.GetObjectIdFromParams(params)

	if err == nil {
		pokemon, err = db.UpdatePokemon(objectId, pokemon)
	}

	return pokemon, err
}

/*
DeletePokemon deletes Pokemon information in the database
*/
func DeletePokemon(params map[string]string) (model.Pokemon, error) {
	var pokemon model.Pokemon
	objectId, err := utils.GetObjectIdFromParams(params)

	if err == nil {
		pokemon, err = db.DeletePokemon(objectId)
	}

	return pokemon, err
}
