package main

import (
	"fmt"

	"first/repository"
	"first/service"
)

func main() {
	pokemonRepository, err := repository.NewPokemonRepository()
	if err != nil {
		panic(err)
	}
	pokemons, err := pokemonRepository.GetAll()
	if err != nil {
		panic(err)
	}
	for _, pokemon := range pokemons {
		fmt.Println(pokemon.Id, pokemon.Name)
	}

	pokemon, err := pokemonRepository.GetById(5)

	if err != nil {
		panic(err)
	}
	fmt.Println(pokemon.Id, pokemon.Name)

	pokemonService, err := service.NewPokemonService()

	if err != nil {
		panic(err)
	}
	pokemonService.Saludo()
}
