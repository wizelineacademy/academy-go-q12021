package models

import (
	"fmt"

	"github.com/dannegm/academy-go-q12021/constants"
	"github.com/dannegm/academy-go-q12021/helpers"
)

// Pokedex stores a list of Pokemon
type Pokedex []Pokemon

// PokedexFromFile return a list of Pokemon
func PokedexFromFile() (pokes Pokedex, err error) {
	rows, err := helpers.ReadFile(fmt.Sprintf("%s/pokemon.csv", constants.AssetsPath))

	if err != nil {
		return
	}

	pokes = Pokedex{}

	for _, row := range rows[1:167] {
		poke, _ := PokemonFromString(row)
		pokes = append(pokes, poke)
	}

	return
}
