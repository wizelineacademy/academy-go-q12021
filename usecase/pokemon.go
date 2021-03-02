package usecase

import (
	"net/http"

	"pokeapi/model"
	csvservice "pokeapi/service/csv"
	httpservice "pokeapi/service/http"
)

const pathFile = "./csv/pokemon.csv"

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
	f, err := us.csvService.Open(pathFile)

	if err != nil {
		return nil, &model.Error{
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
		}
	}
	return us.csvService.GetPokemons(f)
}

func (us *PokemonUsecase) GetPokemon(pokemonId int) (model.Pokemon, *model.Error) {
	f, err := us.csvService.Open(pathFile)

	if err != nil {
		return model.Pokemon{}, &model.Error{
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
		}
	}
	return us.csvService.GetPokemon(pokemonId, f)
}

func (us *PokemonUsecase) GetPokemonsFromExternalAPI() (*[]model.SinglePokeExternal, *model.Error) {
	f, _ := us.csvService.Open(pathFile) //Read only
	lines, _ := us.csvService.ReadAllLines(f)
	fileOpenAndWrite, _ := us.csvService.OpenAndWrite(pathFile) // Write

	newPokemons, httpErr := us.httpService.GetPokemons()

	if httpErr != nil {
		return nil, httpErr
	}

	us.csvService.AddLine(fileOpenAndWrite, lines, &newPokemons)
	return &newPokemons, nil
}
