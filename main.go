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

	pokeservice, err := poke.NewPokeService("poke/pokemon.csv")
	if err != nil {
		log.Fatalln(err)
	}

	ctx := context.Background()
	s := api.NewPokeAPI(ctx, *port, pokeservice)
	log.Printf("POKEAPI available at /getPoke?id using port: %d", port)
	log.Fatalln(s.StartServer())
}
