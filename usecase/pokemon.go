package usecase

import (
	"pokeapi/model"
	"pokeapi/service"
)

type Usecase struct {
	csvService service.NewCsvService
}

type IUsecase interface {
	GetPokemons() ([]model.Pokemon, *model.Error)
	GetPokemon(pokemonId int) (model.Pokemon, *model.Error)
	GetPokemonsFromExternalAPI(newPokemons *[]model.SinglePokeExternal)
}

func New(s service.NewCsvService) *Usecase {
	return &Usecase{s}
}

func (us *Usecase) GetPokemons() ([]model.Pokemon, *model.Error) {
	return us.csvService.GetPokemons()
}

func (us *Usecase) GetPokemon(pokemonId int) (model.Pokemon, *model.Error) {
	return us.csvService.GetPokemon(pokemonId)
}

func (us *Usecase) GetPokemonsFromExternalAPI(newPokes *[]model.SinglePokeExternal) {
	us.csvService.AddLineCsv(newPokes)
}
