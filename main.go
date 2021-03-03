package main

import (
	"context"
	"flag"
	"log"

	"github.com/wizelineacademy/academy-go-q12021/api"
	"github.com/wizelineacademy/academy-go-q12021/poke"
)

func main() {

	port := flag.Int("port", 8080, "the port number")
	flag.Parse()

	pokeservice := poke.NewPokeService("poke/pokemon.csv")
	// pokes := pokeservice.PokemonList
	// fmt.Printf("%+v", pokes["6"])
	ctx := context.Background()
	s := api.NewPokeApi(ctx, *port, pokeservice)
	log.Println("POKEAPI available at /getPoke?id")
	s.StartServer()
}
