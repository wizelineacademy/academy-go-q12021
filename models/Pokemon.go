package models

import (
	"strconv"
	"strings"
)

// Pokemon represet the basic Pokemon model
type Pokemon struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	TypeOne        string `json:"typeOne"`
	TypeTwo        string `json:"typeTwo"`
	Total          int    `json:"total"`
	HealthPoints   int    `json:"healthPoints"`
	Attack         int    `json:"attack"`
	Defense        int    `json:"defense"`
	SpecialAttack  int    `json:"specialAttack"`
	SpecialDefense int    `json:"specialDefense"`
	Speed          int    `json:"speed"`
	Generation     int    `json:"generation"`
	Legendary      bool   `json:"legendary"`
}

// PokemonFromString return a Pokemon from a string
func PokemonFromString(csvRow string) (p Pokemon, err error) {
	pokeRow := strings.Split(string(csvRow), ",")

	id, err := strconv.Atoi(pokeRow[0])
	total, err := strconv.Atoi(pokeRow[4])
	healthPoints, err := strconv.Atoi(pokeRow[5])
	attack, err := strconv.Atoi(pokeRow[6])
	defense, err := strconv.Atoi(pokeRow[7])
	specialAttack, err := strconv.Atoi(pokeRow[8])
	specialDefense, err := strconv.Atoi(pokeRow[9])
	speed, err := strconv.Atoi(pokeRow[10])
	generation, err := strconv.Atoi(pokeRow[11])
	legendary, err := strconv.ParseBool(pokeRow[12])

	p = Pokemon{
		ID:             id,
		Name:           pokeRow[1],
		TypeOne:        pokeRow[2],
		TypeTwo:        pokeRow[3],
		Total:          total,
		HealthPoints:   healthPoints,
		Attack:         attack,
		Defense:        defense,
		SpecialAttack:  specialAttack,
		SpecialDefense: specialDefense,
		Speed:          speed,
		Generation:     generation,
		Legendary:      legendary,
	}
	return
}
