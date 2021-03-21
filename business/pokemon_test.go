package business

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/wizelineacademy/academy-go-q12021/model"
	mocksrepo "github.com/wizelineacademy/academy-go-q12021/repository/mocks"
	mocksservice "github.com/wizelineacademy/academy-go-q12021/service/mocks"
)

func TestGetAllPokemons(t *testing.T) {

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := mocksrepo.NewMockIPokemonRepository(mockCtrl)

	testService, err := NewPokemonBusiness(mockRepo, nil)
	if err != nil {
		panic(err)
	}

	mockRepo.EXPECT().GetAll().Return([]model.Pokemon{
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
	}, nil).Times(1)

	pokemons, _ := testService.GetAll()

	assert.NotNil(t, pokemons)

	mockRepo.EXPECT().GetAll().Return(nil, errors.New("There was an error")).Times(1)

	listPokemons, err := testService.GetAll()

	assert.NotNil(t, err)
	assert.Nil(t, listPokemons)

}

func TestGetByID(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := mocksrepo.NewMockIPokemonRepository(mockCtrl)

	testBusiness, err := NewPokemonBusiness(mockRepo, nil)
	if err != nil {
		panic(err)
	}

	mockRepo.EXPECT().GetByID(gomock.Eq(1)).Return(
		&model.Pokemon{
			Id:             1,
			Name:           "Bulbasaur",
			Height:         10,
			Weight:         20,
			BaseExperience: 100,
			PrimaryType:    "Plant",
			SecondaryType:  "Poison",
		}, nil).Times(1)

	pokemon, err := testBusiness.GetByID(1)

	assert.NotNil(t, &pokemon)
	assert.Nil(t, err)

	mockRepo.EXPECT().GetByID(gomock.Eq(1)).Return(nil, nil).Times(1)

	pokemon2, err := testBusiness.GetByID(1)

	assert.Nil(t, err)
	assert.Nil(t, pokemon2)
}

func TestStoreByID(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockService := mocksservice.NewMockIExternalPokemonAPI(mockCtrl)

	testBusiness, err := NewPokemonBusiness(nil, mockService)
	if err != nil {
		panic(err)
	}

	mockService.EXPECT().GetPokemonFromAPI(1).Return(nil, errors.New("Error")).Times(1)

	pokemon, err := testBusiness.StoreByID(1)

	assert.Nil(t, pokemon)
	assert.NotNil(t, err)

	mockRepo := mocksrepo.NewMockIPokemonRepository(mockCtrl)

	pokemonAPI := model.PokemonAPI{
		Id:             1,
		Name:           "Bulbasaur",
		Height:         10,
		Weight:         20,
		BaseExperience: 100,
		Types: []model.TypeSlot{
			model.TypeSlot{Type: model.Type{Name: "Plant"}},
			model.TypeSlot{Type: model.Type{Name: "Poison"}},
		},
	}

	mockService.EXPECT().GetPokemonFromAPI(gomock.Any()).Return(&pokemonAPI, nil).Times(1)
	mockRepo.EXPECT().StoreToCSV(gomock.Any()).Return(nil, errors.New("Error storing data")).Times(1)

	pokemon2, err := testBusiness.StoreByID(1)

	assert.Nil(t, pokemon2)
	assert.NotNil(t, err)
}
