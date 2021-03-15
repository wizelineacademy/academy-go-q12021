package usecase

import (
	"pokeapi/model"
	csvservice "pokeapi/service/csv"
	httpservice "pokeapi/service/http"
)

type PokemonUsecase struct {
	csvService  csvservice.NewCsvService
	httpService httpservice.NewHttpService
}

type NewPokemonUsecase interface {
	GetPokemons() ([]model.Pokemon, *model.Error)
	GetPokemon(pokemonId int) (model.Pokemon, *model.Error)
	GetPokemonsFromExternalAPI() (*[]model.SinglePokeExternal, *model.Error)
}

func New(s csvservice.NewCsvService, h httpservice.NewHttpService) *PokemonUsecase {
	return &PokemonUsecase{s, h}
}

func (us *PokemonUsecase) GetPokemons() ([]model.Pokemon, *model.Error) {
	return us.csvService.GetPokemons()
}

func (us *PokemonUsecase) GetPokemon(pokemonId int) (model.Pokemon, *model.Error) {
	return us.csvService.GetPokemon(pokemonId)
}

func (us *PokemonUsecase) GetPokemonsFromExternalAPI() (*[]model.SinglePokeExternal, *model.Error) {
	newPokemons, err := us.httpService.GetPokemons()

	if err != nil {
		return nil, err
	}

	errorCsv := us.csvService.SavePokemons(&newPokemons)

	if errorCsv != nil {
		return nil, errorCsv
	}

	return &newPokemons, nil
}
