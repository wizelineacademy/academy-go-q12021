package pokemon

import (
	"bootcamp/domain/model"
	"bootcamp/service/db"
	"gopkg.in/mgo.v2/bson"
)

func GetPokemon() (model.PokemonList, error) {
	pokemonList, err := db.GetPokemon()
	return pokemonList, err
}

func GetPokemonById(objectId bson.ObjectId) (model.Pokemon, error) {
	pokemon, err := db.GetPokemonById(objectId)
	return pokemon, err
}