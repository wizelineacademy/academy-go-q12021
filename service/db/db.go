package db

import (
	"errors"
	"bootcamp/domain/model"
	"bootcamp/utils"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"	
)

func getSession() *mgo.Collection {
	mongoUri, _ := utils.GetEnvVar("MONGO_URL")
	databaseName, _ := utils.GetEnvVar("DB_NAME")
	collectionName, _ := utils.GetEnvVar("COLLECTION")

	session, err := mgo.Dial(mongoUri)

	if err != nil {
		panic(err)
	}

	return session.DB(databaseName).C(collectionName)
}

/*
GetPokemon finds all Pokemon in the mongo collection
*/
func GetPokemon() (model.PokemonList, error) {
	var pokemonList model.PokemonList
	err := getSession().Find(nil).Sort("_id").All(&pokemonList)
	return pokemonList, err
}

/*
GetPokemonById finds a Pokemon by id in the mongo collection
*/
func GetPokemonById(objectId bson.ObjectId) (model.Pokemon, error) {
	var pokemon model.Pokemon
	err := getSession().FindId(objectId).One(&pokemon)
	return pokemon, err
}

/*
AddPokemon inserts a Pokemon in the mongo collection
*/
func AddPokemon(pokemon model.Pokemon) (model.Pokemon, error) {
	err := getSession().Insert(pokemon)
	return pokemon, err
}

/*
UpdatePokemon updates a Pokemon in the mongo collection
*/
func UpdatePokemon(objectId bson.ObjectId, pokemon model.Pokemon) (model.Pokemon, error) {
	document := bson.M{"_id": objectId}
	change := bson.M{"$set":pokemon}
	err := getSession().Update(document, change)
	return pokemon, err
}

/*
DeletePokemon deletes a Pokemon in the mongo collection
*/
func DeletePokemon(objectId bson.ObjectId) (model.Pokemon, error) {
	var pokemon model.Pokemon
	pokemon, err := GetPokemonById(objectId)

	if err != nil {
		return pokemon, errors.New("No pokemon to delete")
	}

	err = getSession().RemoveId(objectId)
	return pokemon, err
}