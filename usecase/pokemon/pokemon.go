package pokemon

import (
	"bootcamp/utils"
	"bootcamp/domain/model"
	"bootcamp/service/db"
	"gopkg.in/mgo.v2/bson"
	"io"
)

func GetPokemon() (model.PokemonList, error) {
	pokemonList, err := db.GetPokemon()
	return pokemonList, err
}

func GetPokemonById(objectId bson.ObjectId) (model.Pokemon, error) {
	pokemon, err := db.GetPokemonById(objectId)
	return pokemon, err
}

func AddPokemon(reader io.ReadCloser) (model.Pokemon, error) {
	pokemon, err := utils.GetPokemonFromReader(reader)
	
	if err == nil {
		pokemon, err = db.AddPokemon(pokemon)
	}

	return pokemon, err
}

func UpdatePokemon(objectId bson.ObjectId, reader io.ReadCloser) (model.Pokemon, error) {
	pokemon, err := utils.GetPokemonFromReader(reader)

	if err == nil {
		pokemon, err = db.UpdatePokemon(objectId, pokemon)
	}

	return pokemon, err
}

func DeletePokemon(objectId bson.ObjectId) (model.Pokemon, error) {
	pokemon, err := db.DeletePokemon(objectId)
	return pokemon, err
}
