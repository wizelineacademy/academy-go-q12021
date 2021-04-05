package usecases

import "github.com/cesararredondow/academy-go-q12021/models"

// GetPokemons this function make the request to csv file to get the pokemons
func (u *UseCase) GetPokemons() ([]*models.Pokemon, error) {
	resp, err := u.service.GetPokemons()
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// GetPokemon this function make the request to csv file to get the pokemon by id
func (u *UseCase) GetPokemon(id string) (*models.Pokemon, error) {
	resp, err := u.service.GetPokemon(id)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
