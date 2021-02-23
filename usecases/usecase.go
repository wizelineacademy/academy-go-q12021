package usecases

import (
	"pokeapi/models"
	"pokeapi/services"
)

type Usecase struct {
	service     services.IService
	httpService services.IHttp
}

type IUsecase interface {
	GetPokemons() []models.Pokemon
	GetPokemon(pokemonId int) models.Pokemon
	GetPokemonsFromExternalAPI()
}

func NewUseCase(s services.IService, hs services.IHttp) *Usecase {
	return &Usecase{s, hs}
}

func (us *Usecase) GetPokemons() []models.Pokemon {
	return us.service.GetPokemons()
}

func (us *Usecase) GetPokemon(pokemonId int) models.Pokemon {
	return us.service.GetPokemon(pokemonId)
}

func (us *Usecase) GetPokemonsFromExternalAPI() {
	us.httpService.GetPokemonsFromExternalAPI()
}
