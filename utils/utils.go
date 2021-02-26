package utils

import (
	"bootcamp/domain/model"
	"log"
	"io"
	"os"
	"encoding/json"
	"strconv"
	"encoding/csv"
	"github.com/spf13/viper"
	"gopkg.in/mgo.v2/bson"
	"errors"
)

func GetEnvVar(key string) string {
  viper.SetConfigFile(".env")
  err := viper.ReadInConfig()

  if err != nil {
    log.Fatalf("Error while reading config file %s", err)
  }

  value, ok := viper.Get(key).(string)

  if !ok {
    log.Fatalf("Invalid type assertion")
  }

  return value
}

func GetObjectIdFromParams(params map[string]string) (bson.ObjectId, error) {
	var objectId bson.ObjectId
	id := params["id"]

	if id == "" || !bson.IsObjectIdHex(id) {
		return objectId,	errors.New("Invalid id provided")
	}

	objectId = bson.ObjectIdHex(id)
	return objectId, nil
}

func GetPokemonFromReader(reader io.ReadCloser) (model.Pokemon, error) {
	var tempPokemon model.Pokemon
	decoder := json.NewDecoder(reader)
	err := decoder.Decode(&tempPokemon)

	if err == nil {
		defer reader.Close()		
	}

	return tempPokemon, err
}

func ReadCSV() (model.PokemonList, error) {
	var pokemonList model.PokemonList
	recordFile, err := os.Open("assets/pokemon.csv")

	if err == nil {
		reader := csv.NewReader(recordFile)
		records, err := reader.ReadAll()

		if err == nil {
			for _, pokemon := range records {
				id, err := strconv.Atoi(pokemon[0])
		
				if err != nil {
					 return pokemonList, errors.New("Cannot get id from row")
				}
				
				pkNumber, err := strconv.Atoi(pokemon[1])
		
				if err != nil {
					return pokemonList, errors.New("Cannot get pokedex number from row")
				}
		
				pk := model.Pokemon{Id:id, PokedexNumber: pkNumber, Name:pokemon[2], Types:pokemon[3], Region:pokemon[4]}
				pokemonList = append(pokemonList, pk)
			}
		
			err = recordFile.Close()
		
			if err != nil {
				return pokemonList, errors.New("Error while closing file")
			}
		
			return pokemonList, nil
		}
		return nil, err
	}
	return nil, err
}