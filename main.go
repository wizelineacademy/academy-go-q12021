package main

import (
	"fmt"
	"log"

	"github.com/adantop/golang-bootcamp-2020/db"
	"github.com/adantop/golang-bootcamp-2020/fs"
	"github.com/adantop/golang-bootcamp-2020/pokemon"
	"github.com/docopt/docopt-go"
)

func main() {
	usage := `Pokedex

Usage:
  pokedex csv <csvfile> <PokemonName>
  pokedex sqlite3 <dbfile> <PokemonName>
  pokedex postgres <PokemonName>
  pokedex -h | --help

Options:
  -h --help     Show this screen.`

	args, _ := docopt.ParseDoc(usage)

	var ds *pokemon.DataSource

	switch {
	case args["csv"]:
		ds = &fs.DS
		csvfile, _ := args.String("<csvfile>")
		fs.UseCSV(csvfile)
	case args["postgres"]:
		ds = &db.DS
		db.UsePostgreSQL()
	case args["sqlite"]:
		fallthrough
	default:
		ds = &db.DS
		dbfile, _ := args.String("<dbfile>")
		db.UseSQLite3(dbfile)
	}
	defer (*ds).Close()

	name, _ := args.String("<PokemonName>")
	pokemon, err := (*ds).GetPokemonByName(name)

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(pokemon.Describe())
}
