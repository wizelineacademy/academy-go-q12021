package domain

import (
	"fmt"
)

// Pokemon struct definition
type Pokemon struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
}

// Pokemon api struct definition
type PokemonApiElement struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	URL       string    `json:"url"`
}

// Pokemon api struct definition
type PokemonApiRequest struct {
	Count          int    `json:"count"`
	Next           string    `json:"next"`
	Previous       string    `json:"previous"`
	Results        []PokemonApiElement
}

// NewPokemon returns a new Pokemon struct
func NewPokemon(id int, name string) *Pokemon {
	return &Pokemon{
		ID:   id,
		Name: name,
	}
}

// Info - displays Pokemon's information
func (*pokemon Pokemon) Info() string {
	return fmt.Sprintf("Pokemons, name '%s'\n", pokemon.name)
}
