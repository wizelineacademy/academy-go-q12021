package utils

import (
	"models"
	"fmt"
	"log"
	"os"
	"strconv"
	"encoding/csv"
	"github.com/spf13/viper"
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

func ReadCSV() models.PokemonList {
	var pokeList models.PokemonList

	recordFile, err := os.Open("pokemon.csv")
	if err != nil {
		fmt.Println("Error while openning file:", err)
		return pokeList
	}

	reader := csv.NewReader(recordFile)
	records, err := reader.ReadAll()

	if err != nil {
		fmt.Println("Error retriving all rows:", err)
		return pokeList
	}

	for _, pokemon := range records {
		id, err := strconv.Atoi(pokemon[0])

		if err != nil {
			 fmt.Println("Cannot get id from row")
		}

		poke := models.Pokemon{Id:id, Name:pokemon[1], Types:pokemon[2], Region:pokemon[3]}
		pokeList = append(pokeList, poke)
	}

	err = recordFile.Close()

	if err != nil {
		fmt.Println("Error while closing file:", err)
		return pokeList
	}

	return pokeList
}