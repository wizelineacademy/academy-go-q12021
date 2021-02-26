package pokemon

import (
	"io"
	"bootcamp/domain/model"
	"bootcamp/service/db"
	"gopkg.in/mgo.v2/bson"
	"encoding/json"
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
	var tempPokemon model.Pokemon
	decoder := json.NewDecoder(reader)
	err := decoder.Decode(&tempPokemon)

	if err != nil {
		return tempPokemon, err
	}

	defer reader.Close()
	pokemon, err := db.AddPokemon(tempPokemon)
	return pokemon, err
}

func UpdatePokemon(objectId bson.ObjectId, reader io.ReadCloser) (model.Pokemon, error) {
	var tempPokemon model.Pokemon
	decoder := json.NewDecoder(reader)
	err := decoder.Decode(&tempPokemon)

	if err != nil {
		return tempPokemon, err
	}

	defer reader.Close()
	pokemon, err := db.UpdatePokemon(objectId, tempPokemon)
	return pokemon, err
}
