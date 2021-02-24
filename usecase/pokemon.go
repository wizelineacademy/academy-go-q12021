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
	GetPokemons() ([]model.Pokemon, *model.Error)
	GetPokemon(pokemonId int) (model.Pokemon, *model.Error)
	GetPokemonsFromExternalAPI() *model.Error
}

func New(s service.NewCsvService, hs service.NewHttpService) *Usecase {
	return &Usecase{s, hs}
}

func (us *Usecase) GetPokemons() ([]model.Pokemon, *model.Error) {
	return us.csvService.GetPokemons()
}

func (us *Usecase) GetPokemon(pokemonId int) (model.Pokemon, *model.Error) {
	return us.csvService.GetPokemon(pokemonId)
}

func (us *Usecase) GetPokemonsFromExternalAPI() *model.Error {
	return us.httpService.GetPokemonsFromExternalAPI()
}
