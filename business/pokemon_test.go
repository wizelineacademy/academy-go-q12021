package business

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/wizelineacademy/academy-go-q12021/model"
)

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) GetAll() ([]model.Pokemon, error) {
	args := m.Called()
	result := args.Get(0)
	return result.([]model.Pokemon), args.Error(1)
}
func (m *MockRepository) OpenFile() (*os.File, error) {
	args := m.Called()
	result := args.Get(0)
	return result.(*os.File), args.Error(1)
}
func (m *MockRepository) GetByID(id int) (*model.Pokemon, error) {
	args := m.Called()
	result := args.Get(0)
	return result.(*model.Pokemon), args.Error(1)
}
func (m *MockRepository) StoreToCSV(pokemonAPI model.PokemonAPI) (*model.Pokemon, error) {
	args := m.Called()
	result := args.Get(0)
	return result.(*model.Pokemon), args.Error(1)
}
func (m *MockRepository) GetCSVDataInMemory() (map[int]model.Pokemon, error) {
	args := m.Called()
	result := args.Get(0)
	return result.(map[int]model.Pokemon), args.Error(1)
}
func (m *MockRepository) CloseFile(file *os.File) {
}

func TestGetAllPokemonsFromRepository(t *testing.T) {
	mockRepo := new(MockRepository)
	testService, err := NewPokemonBusiness(mockRepo, nil)
	if err != nil {
		panic(err)
	}

	mockRepo.On("GetAll").Return([]model.Pokemon{
		model.Pokemon{
			Id:             3,
			Name:           "Bulbasaur",
			Height:         10,
			Weight:         20,
			BaseExperience: 100,
			PrimaryType:    "Plant",
			SecondaryType:  "Poison",
		},
		model.Pokemon{
			Id:             10,
			Name:           "Charmander",
			Height:         10,
			Weight:         20,
			BaseExperience: 100,
			PrimaryType:    "Plant",
			SecondaryType:  "Poison",
		},
	}, nil)

	pokemons, _ := testService.GetAll()

	assert.NotNil(t, pokemons)
}
