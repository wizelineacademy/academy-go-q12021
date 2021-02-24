package usecase

import (
	"pokeapi/model"
	"pokeapi/service"
)

type Usecase struct {
	csvService  service.NewCsvService
	httpService service.NewHttpService
}

type IUsecase interface {
	GetPokemons() []model.Pokemon
	GetPokemon(pokemonId int) model.Pokemon
	GetPokemonsFromExternalAPI()
}

func New(s service.NewCsvService, hs service.NewHttpService) *Usecase {
	return &Usecase{s, hs}
}

func (us *Usecase) GetPokemons() []model.Pokemon {
	return us.csvService.GetPokemons()
}

func (us *Usecase) GetPokemon(pokemonId int) model.Pokemon {
	return us.csvService.GetPokemon(pokemonId)
}

func (us *Usecase) GetPokemonsFromExternalAPI() {
	us.httpService.GetPokemonsFromExternalAPI()
}
