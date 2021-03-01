package usecase

import (
	"net/http"
	"pokeapi/model"
	"pokeapi/service"
)

type Usecase struct {
	csvService service.NewCsvService
}

type IUsecase interface {
	GetPokemons() ([]model.Pokemon, *model.Error)
	GetPokemon(pokemonId int) (model.Pokemon, *model.Error)
	GetPokemonsFromExternalAPI(newPokemons *[]model.SinglePokeExternal) *model.Error
}

func New(s service.NewCsvService) *Usecase {
	return &Usecase{s}
}

func (us *Usecase) GetPokemons() ([]model.Pokemon, *model.Error) {
	const pathFile = "./csv/pokemon.csv"
	f, err := us.csvService.Open(pathFile)

	if err != nil {
		return nil, &model.Error{
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
		}
	}
	return us.csvService.GetPokemons(f)
}

func (us *Usecase) GetPokemon(pokemonId int) (model.Pokemon, *model.Error) {
	const pathFile = "./csv/pokemon.csv"
	f, err := us.csvService.Open(pathFile)

	if err != nil {
		return model.Pokemon{}, &model.Error{
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
		}
	}
	return us.csvService.GetPokemon(pokemonId, f)
}

func (us *Usecase) GetPokemonsFromExternalAPI(newPokes *[]model.SinglePokeExternal) *model.Error {
	const pathFile = "./csv/pokemon.csv"
	f, _ := us.csvService.Open(pathFile) //Read only
	lines, _ := us.csvService.ReadAllLines(f)
	f, err := us.csvService.OpenAndWrite(pathFile) // Write

	if err != nil {
		return &model.Error{
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
		}
	}

	return us.csvService.AddLine(f, lines, newPokes)
}
