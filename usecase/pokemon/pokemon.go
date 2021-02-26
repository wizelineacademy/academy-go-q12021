package pokemon

import (
	"io"
	"bootcamp/domain/model"
	"bootcamp/service/db"
	"gopkg.in/mgo.v2/bson"
	"encoding/json"
)

func getPokemonFromReader(reader io.ReadCloser) (model.Pokemon, error) {
	var tempPokemon model.Pokemon
	decoder := json.NewDecoder(reader)
	err := decoder.Decode(&tempPokemon)

	if err == nil {
		defer reader.Close()		
	}

	return tempPokemon, err
}

func GetPokemon() (model.PokemonList, error) {
	pokemonList, err := db.GetPokemon()
	return pokemonList, err
}

func GetPokemonById(objectId bson.ObjectId) (model.Pokemon, error) {
	pokemon, err := db.GetPokemonById(objectId)
	return pokemon, err
}

func AddPokemon(reader io.ReadCloser) (model.Pokemon, error) {
	pokemon, err := getPokemonFromReader(reader)
	
	if err == nil {
		pokemon, err = db.AddPokemon(pokemon)
	}

	return pokemon, err
}

func UpdatePokemon(objectId bson.ObjectId, reader io.ReadCloser) (model.Pokemon, error) {
	pokemon, err := getPokemonFromReader(reader)

	if err == nil {
		pokemon, err = db.UpdatePokemon(objectId, pokemon)
	}

	return pokemon, err
}

func DeletePokemon(objectId bson.ObjectId) (model.Pokemon, error) {
	pokemon, err := db.DeletePokemon(objectId)
	return pokemon, err
}
