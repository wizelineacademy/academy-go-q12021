package usecases

import (
	"strconv"

	"github.com/cesararredondow/academy-go-q12021/models"
)

//GetPokemonsConcurrency is the usecase to get the information
func (u *UseCase) GetPokemonsConcurrency(odd string, items string, workers string) ([]*models.Pokemon, error) {
	pokemons, err := u.service.GetPokemons()
	if err != nil {
		return nil, err
	}

	itemsNumber, err := strconv.Atoi(items)
	if err != nil {
		return nil, err
	}

	workersNumber, err := strconv.Atoi(workers)
	if err != nil {
		return nil, err
	}

	booleanOdd, err := strconv.ParseBool(odd)
	if err != nil {
		return nil, err
	}

	resp, err := u.service.GetRegistries(booleanOdd, itemsNumber, workersNumber, pokemons)

	return resp, err
}
