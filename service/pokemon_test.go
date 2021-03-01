package service_test

import (
	"pokeapi/model"
	"pokeapi/service"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOpen(t *testing.T) {
	pathFile := "../csv/pokemon.csv"
	csvService := service.New()
	f, _ := csvService.Open(pathFile)

	assert.NotNil(t, f)
}

func TestOpenError(t *testing.T) {
	pathFile := "./csv/pokemon2.csv"
	csvService := service.New()
	_, err := csvService.Open(pathFile)

	assert.Equal(t, err.Error(), "There was an error opening the file")
}

func TestRead(t *testing.T) {
	pathFile := "../csv/pokemon.csv"
	csvService := service.New()
	f, _ := csvService.Open(pathFile)
	pokemons, _ := service.Read(f)
	mockPokemons := []model.Pokemon{
		{ID: 1, Name: "greninja", URL: "https://pokeapi.co/api/v2/pokemon/658/"},
		{ID: 2, Name: "ursaring", URL: "https://pokeapi.co/api/v2/pokemon/217/"},
		{ID: 3, Name: "arcanine", URL: "https://pokeapi.co/api/v2/pokemon/59/"},
		{ID: 4, Name: "gengar", URL: "https://pokeapi.co/api/v2/pokemon/94/"},
		{ID: 5, Name: "porygon", URL: "https://pokeapi.co/api/v2/pokemon/137/"},
	}

	assert.Equal(t, mockPokemons, pokemons)
}

func TestAddLine(t *testing.T) {
	pathFile := "../csv/pokemon.csv"

	csvService := service.New()
	fWrite, _ := csvService.OpenAndWrite(pathFile)
	f, _ := csvService.Open(pathFile)
	lines, _ := csvService.ReadAllLines(f)

	mockPokemons := []model.SinglePokeExternal{
		{Name: "delcatty", URL: "https://pokeapi.co/api/v2/pokemon/301/"},
		{Name: "sableye", URL: "https://pokeapi.co/api/v2/pokemon/302/"},
		{Name: "mawile", URL: "https://pokeapi.co/api/v2/pokemon/303/"},
		{Name: "aron", URL: "https://pokeapi.co/api/v2/pokemon/304/"},
		{Name: "lairon", URL: "https://pokeapi.co/api/v2/pokemon/305/"},
	}
	err := csvService.AddLine(fWrite, lines, &mockPokemons)

	assert.Nil(t, err)
}
