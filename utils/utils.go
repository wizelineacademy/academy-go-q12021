package utils

import (
	"model"
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

func ReadCSV() model.PokemonList {
	var pokeList model.PokemonList

	recordFile, err := os.Open("assets/pokemon.csv")
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
		
		pkNumber, err2 := strconv.Atoi(pokemon[1])

		if err2 != nil {
			 fmt.Println("Cannot get pokedex number from row")
		}

		poke := model.Pokemon{Id:id, PokedexNumber: pkNumber, Name:pokemon[2], Types:pokemon[3], Region:pokemon[4]}
		pokeList = append(pokeList, poke)
	}

	err = recordFile.Close()

	if err != nil {
		fmt.Println("Error while closing file:", err)
		return pokeList
	}

	return pokeList
}