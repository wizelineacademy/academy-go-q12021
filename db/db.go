package db

import (
	"github.com/adantop/golang-bootcamp-2020/pokemon"
)

var (
	// DS Is the pokemon Datasource
	DS                 pokemon.DataSource
	queryPokemonByName = "SELECT * FROM pokemon WHERE name = $1"
)
