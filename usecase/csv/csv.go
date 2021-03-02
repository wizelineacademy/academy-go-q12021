package csv

import (
	"errors"
	"net/http"
	"strconv"
	"bootcamp/domain/model"
	"bootcamp/utils"
	"github.com/gorilla/mux"
)

/*
GetPokemon returns a Pokemon list read from a CSV file
If query params exists or /{id}, filter the Pokemon
Otherwise, returns the list
*/
func GetPokemon(r *http.Request) (model.PokemonList, error) {
	pokemonList, err := utils.ReadCSV()
	var pokemonSubset model.PokemonList
	params := mux.Vars(r)

	if err == nil {
		id:= params["id"]

		if id != "" {
			index, _ := strconv.Atoi(id)
			pokemonId := index - 1
	
			if pokemonId <= len(pokemonList) - 1 {
				pokemonSubset = append(pokemonSubset, pokemonList[pokemonId])
				return pokemonSubset, nil
			}

			err = errors.New("Invalid index")
			return nil, err
		}
	}

	queryParams := r.URL.Query()

	if len(queryParams) > 0 {
		pokemon:= utils.GetPokemonByKey(queryParams, pokemonList)
		var pokemonSubset model.PokemonList
		pokemonSubset = append(pokemonSubset, pokemon)
		return pokemonSubset, nil
	}

	return pokemonList, err
}