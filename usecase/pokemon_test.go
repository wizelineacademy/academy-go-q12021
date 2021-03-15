package usecase

import (
	"reflect"
	"testing"

	"pokeapi/model"
	csvservice "pokeapi/service/csv"
	httpservice "pokeapi/service/http"
	"pokeapi/service/mock"

	"github.com/golang/mock/gomock"
)

func TestPokemonUsecase_GetPokemons(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCsvService := mock.NewMockNewCsvService(ctrl)
	mockCsvService.EXPECT().GetPokemons().Return([]model.Pokemon{
		{ID: 1, Name: "greninja", URL: "https://pokeapi.co/api/v2/pokemon/658/"},
		{ID: 2, Name: "ursaring", URL: "https://pokeapi.co/api/v2/pokemon/217/"},
		{ID: 3, Name: "arcanine", URL: "https://pokeapi.co/api/v2/pokemon/59/"},
		{ID: 4, Name: "gengar", URL: "https://pokeapi.co/api/v2/pokemon/94/"},
		{ID: 5, Name: "porygon", URL: "https://pokeapi.co/api/v2/pokemon/137/"},
	}, nil)

	tests := []struct {
		name       string
		csvService csvservice.NewCsvService
		want       []model.Pokemon
		wantErr    *model.Error
	}{
		{
			name:       "succeded Get Pokemons",
			csvService: mockCsvService,
			want: []model.Pokemon{
				{ID: 1, Name: "greninja", URL: "https://pokeapi.co/api/v2/pokemon/658/"},
				{ID: 2, Name: "ursaring", URL: "https://pokeapi.co/api/v2/pokemon/217/"},
				{ID: 3, Name: "arcanine", URL: "https://pokeapi.co/api/v2/pokemon/59/"},
				{ID: 4, Name: "gengar", URL: "https://pokeapi.co/api/v2/pokemon/94/"},
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
			got, gotErr := us.GetPokemons()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PokemonUsecase.GetPokemons() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(gotErr, tt.wantErr) {
				t.Errorf("PokemonUsecase.GetPokemons() got1 = %v, want %v", gotErr, tt.wantErr)
			}
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
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PokemonUsecase.GetPokemon() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(gotErr, tt.wantErr) {
				t.Errorf("PokemonUsecase.GetPokemon() got1 = %v, want %v", gotErr, tt.wantErr)
			}
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
	mockNewPokemons := []model.SinglePokeExternal{
		{Name: "delcatty", URL: "https://pokeapi.co/api/v2/pokemon/301/"},
		{Name: "sableye", URL: "https://pokeapi.co/api/v2/pokemon/302/"},
		{Name: "mawile", URL: "https://pokeapi.co/api/v2/pokemon/303/"},
		{Name: "aron", URL: "https://pokeapi.co/api/v2/pokemon/304/"},
		{Name: "lairon", URL: "https://pokeapi.co/api/v2/pokemon/305/"},
	}

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
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PokemonUsecase.GetPokemonsFromExternalAPI() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(gotErr, tt.wantErr) {
				t.Errorf("PokemonUsecase.GetPokemonsFromExternalAPI() got1 = %v, want %v", gotErr, tt.wantErr)
			}
		})
	}
}
