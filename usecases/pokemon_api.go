package usecases

import "github.com/cesararredondow/academy-go-q12021/models"

//GetPokemons is the usecase to get the information
func (u *UseCase) GetPokemonsFromAPI(quantity string) ([]*models.Pokemon_api, error) {
	resp, err := u.service.GetPokemonsFromAPI(quantity)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

////GetPokemon is the usecase to get the information
func (u *UseCase) GetPokemonFromAPI(pokemonID string) (*models.PokemonResponse, error) {
	resp, err := u.service.GetPokemonFromAPI(pokemonID)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
