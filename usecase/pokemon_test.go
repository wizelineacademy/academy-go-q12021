package usecase

import (
	"pokeapi/model"
	csvservice "pokeapi/service/csv"
	httpservice "pokeapi/service/http"
	"pokeapi/service/mock"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

var pokemonsTest = []model.Pokemon{
	{ID: 1, Name: "greninja", URL: "https://pokeapi.co/api/v2/pokemon/658/"},
	{ID: 2, Name: "ursaring", URL: "https://pokeapi.co/api/v2/pokemon/217/"},
	{ID: 3, Name: "arcanine", URL: "https://pokeapi.co/api/v2/pokemon/59/"},
	{ID: 4, Name: "gengar", URL: "https://pokeapi.co/api/v2/pokemon/94/"},
	{ID: 5, Name: "porygon", URL: "https://pokeapi.co/api/v2/pokemon/137/"},
	{ID: 6, Name: "flareon", URL: "https://pokeapi.co/api/v2/pokemon/136/"},
	{ID: 7, Name: "omanyte", URL: "https://pokeapi.co/api/v2/pokemon/138/"},
	{ID: 8, Name: "frillish", URL: "https://pokeapi.co/api/v2/pokemon/592/"},
	{ID: 9, Name: "cacturne", URL: "https://pokeapi.co/api/v2/pokemon/332/"},
	{ID: 10, Name: "scizor", URL: "https://pokeapi.co/api/v2/pokemon/212/"},
}

var pokemonsFromHttp = []model.SinglePokeExternal{
	{Name: "delcatty", URL: "https://pokeapi.co/api/v2/pokemon/301/"},
	{Name: "sableye", URL: "https://pokeapi.co/api/v2/pokemon/302/"},
	{Name: "mawile", URL: "https://pokeapi.co/api/v2/pokemon/303/"},
	{Name: "aron", URL: "https://pokeapi.co/api/v2/pokemon/304/"},
	{Name: "lairon", URL: "https://pokeapi.co/api/v2/pokemon/305/"},
}

func TestPokemonUsecase_GetPokemons(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCsvService := mock.NewMockNewCsvService(ctrl)
	mockCsvService.EXPECT().GetPokemons().Return(pokemonsTest, nil)

	tests := []struct {
		name       string
		csvService csvservice.NewCsvService
		want       []model.Pokemon
		wantErr    *model.Error
	}{
		{
			name:       "succeded Get Pokemons",
			csvService: mockCsvService,
			want:       pokemonsTest,
			wantErr:    nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			us := &PokemonUsecase{
				csvService: tt.csvService,
			}
			got, gotErr := us.GetPokemons()
			assert.Equal(t, got, tt.want)
			assert.Equal(t, gotErr, tt.wantErr)
		})
	}
}

func TestPokemonUsecase_GetPokemon(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCsvService := mock.NewMockNewCsvService(ctrl)
	mockCsvService.EXPECT().GetPokemon(1).Return(model.Pokemon{
		ID:   1,
		Name: "greninja",
		URL:  "https://pokeapi.co/api/v2/pokemon/658/",
	}, nil)
	mockCsvService.EXPECT().GetPokemon(4).Return(model.Pokemon{
		ID:   4,
		Name: "gengar",
		URL:  "https://pokeapi.co/api/v2/pokemon/94/",
	}, nil)

	tests := []struct {
		name       string
		pokemonId  int
		csvService csvservice.NewCsvService
		want       model.Pokemon
		wantErr    *model.Error
	}{
		{
			name:       "succeded Get Pokemon",
			pokemonId:  1,
			csvService: mockCsvService,
			want: model.Pokemon{
				ID:   1,
				Name: "greninja",
				URL:  "https://pokeapi.co/api/v2/pokemon/658/",
			},
			wantErr: nil,
		},
		{
			name:       "succeded Get Pokemon",
			pokemonId:  4,
			csvService: mockCsvService,
			want: model.Pokemon{
				ID:   4,
				Name: "gengar",
				URL:  "https://pokeapi.co/api/v2/pokemon/94/",
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			us := &PokemonUsecase{
				csvService: tt.csvService,
			}
			got, gotErr := us.GetPokemon(tt.pokemonId)
			assert.Equal(t, got, tt.want)
			assert.Equal(t, gotErr, tt.wantErr)
		})
	}
}

func TestPokemonUsecase_GetPokemonsFromExternalAPI(t *testing.T) {
	// We need to define the controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// We need to get an instance of our MockServices
	mockCsvService := mock.NewMockNewCsvService(ctrl)
	mockHttpService := mock.NewMockNewHttpService(ctrl)

	// We create the mocking new Pokemons array retrieved by the http Service
	mockNewPokemons := pokemonsFromHttp

	// we mock our methods
	mockHttpService.EXPECT().GetPokemons().Return(mockNewPokemons, nil)
	mockCsvService.EXPECT().SavePokemons(&mockNewPokemons).Return(nil)

	tests := []struct {
		name        string
		csvService  csvservice.NewCsvService
		httpService httpservice.NewHttpService
		want        *[]model.SinglePokeExternal
		wantErr     *model.Error
	}{
		{
			name:        "Succeded GetPokemonsFromExternalAPI",
			csvService:  mockCsvService,
			httpService: mockHttpService,
			want:        &mockNewPokemons,
			wantErr:     nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			us := &PokemonUsecase{
				csvService:  tt.csvService,
				httpService: tt.httpService,
			}
			got, gotErr := us.GetPokemonsFromExternalAPI()
			assert.Equal(t, got, tt.want)
			assert.Equal(t, gotErr, tt.wantErr)
		})
	}
}

func TestPokemonUsecase_GetPokemonsConcurrently(t *testing.T) {
	// We need to define the controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// We need to get an instance of our MockServices
	mockCsvService := mock.NewMockNewCsvService(ctrl)

	// we need to mock our methods
	mockCsvService.EXPECT().GetPokemons().Return(pokemonsTest, nil) // we do need to call the function same number of tests we run
	mockCsvService.EXPECT().GetPokemons().Return(pokemonsTest, nil)
	mockCsvService.EXPECT().GetPokemons().Return(pokemonsTest, nil)

	tests := []struct {
		name           string
		csvService     csvservice.NewCsvService
		typeNumber     string
		items          int
		itemsPerWorker int
		want           []model.Pokemon
		wantErr        *model.Error
	}{
		{
			name:           "Succeded Pokemons: odd, items: 1, itemsPerWorker: 1",
			csvService:     mockCsvService,
			typeNumber:     "odd",
			items:          1,
			itemsPerWorker: 1,
			want: []model.Pokemon{
				{
					ID:   1,
					Name: "greninja",
					URL:  "https://pokeapi.co/api/v2/pokemon/658/",
				},
			},
			wantErr: nil,
		},
		{
			name:           "Succeded Pokemons: even, items: 3, itemsPerWorker: 2",
			csvService:     mockCsvService,
			typeNumber:     "even",
			items:          3,
			itemsPerWorker: 2,
			want: []model.Pokemon{
				{ID: 2, Name: "ursaring", URL: "https://pokeapi.co/api/v2/pokemon/217/"},
				{ID: 4, Name: "gengar", URL: "https://pokeapi.co/api/v2/pokemon/94/"},
				{ID: 6, Name: "flareon", URL: "https://pokeapi.co/api/v2/pokemon/136/"},
			},
			wantErr: nil,
		},
		{
			name:           "Succeded Pokemons: even, items: 10, itemsPerWorker: 3",
			csvService:     mockCsvService,
			typeNumber:     "odd",
			items:          10,
			itemsPerWorker: 3,
			want: []model.Pokemon{
				{ID: 1, Name: "greninja", URL: "https://pokeapi.co/api/v2/pokemon/658/"},
				{ID: 3, Name: "arcanine", URL: "https://pokeapi.co/api/v2/pokemon/59/"},
				{ID: 7, Name: "omanyte", URL: "https://pokeapi.co/api/v2/pokemon/138/"},
				{ID: 9, Name: "cacturne", URL: "https://pokeapi.co/api/v2/pokemon/332/"},
				{ID: 5, Name: "porygon", URL: "https://pokeapi.co/api/v2/pokemon/137/"},
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			us := &PokemonUsecase{
				csvService: tt.csvService,
			}
			got, gotErr := us.GetPokemonsConcurrently(tt.typeNumber, tt.items, tt.itemsPerWorker)
			assert.ElementsMatch(t, got, tt.want)
			assert.Equal(t, gotErr, tt.wantErr)
		})
	}
}
