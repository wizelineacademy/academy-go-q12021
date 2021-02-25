package db

import (
	"bootcamp/domain/model"
	"bootcamp/utils"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func getSession() *mgo.Collection {
	mongoUri := utils.GetEnvVar("MONGO_URL")
	databaseName := utils.GetEnvVar("DB_NAME")
	collectionName := utils.GetEnvVar("COLLECTION")

	session, err := mgo.Dial(mongoUri)

	if err != nil {
		panic(err)
	}

	return session.DB(databaseName).C(collectionName)
}

func GetPokemon() (model.PokemonList, error) {
	var pokemonList model.PokemonList
	err := getSession().Find(nil).Sort("_id").All(&pokemonList)
	return pokemonList, err
}

func GetPokemonById(objectId bson.ObjectId) (model.Pokemon, error) {
	var pokemon model.Pokemon
	err := getSession().FindId(objectId).One(&pokemon)
	return pokemon, err
}