package usecase_test

import (
	"pokeapi/mock"
	"pokeapi/service"
	"pokeapi/usecase"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPokemons(t *testing.T) {
	// This test still failing, pokemons
	// is returned as nil
	csvService := service.New()
	usecase := usecase.New(csvService)

	pokemons, _ := usecase.GetPokemons()

	mockPokemons := mock.MockPokemon()
	assert.Equal(t, pokemons, mockPokemons)
}
