package controller

import (
	"net/http"
	"net/http/httptest"
	"pokeapi/model"
	"pokeapi/usecase"
	usecasemock "pokeapi/usecase/mock"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
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

var pokemonsFromHttp = &[]model.SinglePokeExternal{
	{Name: "delcatty", URL: "https://pokeapi.co/api/v2/pokemon/301/"},
	{Name: "sableye", URL: "https://pokeapi.co/api/v2/pokemon/302/"},
	{Name: "mawile", URL: "https://pokeapi.co/api/v2/pokemon/303/"},
	{Name: "aron", URL: "https://pokeapi.co/api/v2/pokemon/304/"},
	{Name: "lairon", URL: "https://pokeapi.co/api/v2/pokemon/305/"},
}

func TestPokemonController_Index(t *testing.T) {

	request := httptest.NewRequest("GET", "/", nil)
	recorder := httptest.NewRecorder()

	tests := []struct {
		name           string
		r              *http.Request
		rr             *httptest.ResponseRecorder
		want           string
		wantStatusCode int
	}{
		{
			name:           "Succeded Index Http Request",
			r:              request,
			rr:             recorder,
			want:           `{ "message": "Welcome to my Poke-API" }`,
			wantStatusCode: http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pc := &PokemonController{
				useCase: nil,
			}
			handler := http.HandlerFunc(pc.Index)
			handler.ServeHTTP(tt.rr, tt.r)

			// Check the status code is what we expect.
			assert.Equal(t, tt.rr.Code, tt.wantStatusCode)

			// Check the response body is what we expect.
			assert.Equal(t, tt.rr.Body.String(), tt.want)
		})
	}
}

func TestPokemonController_GetPokemons(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	pokemons := pokemonsTest

	mockUsecasePokemon := usecasemock.NewMockNewPokemonUsecase(ctrl)
	mockUsecasePokemon.EXPECT().GetPokemons().Return(pokemons, nil)

	request := httptest.NewRequest("GET", "/pokemons", nil)
	recorder := httptest.NewRecorder()

	tests := []struct {
		name           string
		useCase        usecase.NewPokemonUsecase
		r              *http.Request
		rr             *httptest.ResponseRecorder
		want           []model.Pokemon
		wantStatusCode int
	}{
		{
			name:           "succeded Get Pokemons",
			useCase:        mockUsecasePokemon,
			r:              request,
			rr:             recorder,
			want:           pokemons,
			wantStatusCode: http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pc := &PokemonController{
				useCase: tt.useCase,
			}
			handler := http.HandlerFunc(pc.GetPokemons)
			handler.ServeHTTP(tt.rr, tt.r)

			assert.Equal(t, tt.rr.Code, tt.wantStatusCode)

			reflect.DeepEqual(tt.rr.Body, tt.want)
		})
	}
}

func TestPokemonController_GetPokemon(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	pokemon := model.Pokemon{
		ID:   1,
		Name: "greninja",
		URL:  "https://pokeapi.co/api/v2/pokemon/658/",
	}
	pokemonTest := model.Pokemon{
		ID: 3, Name: "arcanine", URL: "https://pokeapi.co/api/v2/pokemon/59/",
	}
	mockUsecasePokemon := usecasemock.NewMockNewPokemonUsecase(ctrl)
	mockUsecasePokemon.EXPECT().GetPokemon(1).Return(pokemon, nil)
	mockUsecasePokemon.EXPECT().GetPokemon(3).Return(pokemonTest, nil)

	request := httptest.NewRequest("GET", "/pokemons/1", nil)
	recorder := httptest.NewRecorder()

	requestSecondTest := httptest.NewRequest("GET", "/pokemons/3", nil)
	recorderSecondTest := httptest.NewRecorder()

	tests := []struct {
		name           string
		useCase        usecase.NewPokemonUsecase
		r              *http.Request
		rr             *httptest.ResponseRecorder
		want           model.Pokemon
		wantStatusCode int
		pokemonId      int
	}{
		{
			name:           "Succeded Get Pokemon Id: 1",
			useCase:        mockUsecasePokemon,
			r:              request,
			rr:             recorder,
			want:           pokemon,
			pokemonId:      1,
			wantStatusCode: http.StatusOK,
		},
		{
			name:           "Succeded Get Pokemon Id: 3",
			useCase:        mockUsecasePokemon,
			r:              requestSecondTest,
			rr:             recorderSecondTest,
			want:           pokemonTest,
			pokemonId:      3,
			wantStatusCode: http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pc := &PokemonController{
				useCase: tt.useCase,
			}
			pc.useCase.GetPokemon(tt.pokemonId)
			handler := http.HandlerFunc(pc.GetPokemon)
			handler(tt.rr, tt.r)

			assert.Equal(t, tt.rr.Code, tt.wantStatusCode)

			reflect.DeepEqual(tt.rr.Body, tt.want)
		})
	}
}

func TestPokemonController_GetPokemonsFromExternalAPI(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	newPokemons := pokemonsFromHttp
	mockUsecasePokemon := usecasemock.NewMockNewPokemonUsecase(ctrl)
	mockUsecasePokemon.EXPECT().GetPokemonsFromExternalAPI().Return(newPokemons, nil)

	request := httptest.NewRequest("GET", "/pokemons/external", nil)
	recorder := httptest.NewRecorder()

	tests := []struct {
		name           string
		useCase        usecase.NewPokemonUsecase
		r              *http.Request
		rr             *httptest.ResponseRecorder
		want           *[]model.SinglePokeExternal
		wantStatusCode int
	}{
		{
			name:           "Succeded Get Pokemon From External API",
			useCase:        mockUsecasePokemon,
			r:              request,
			rr:             recorder,
			want:           newPokemons,
			wantStatusCode: http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pc := &PokemonController{
				useCase: tt.useCase,
			}
			pc.GetPokemonsFromExternalAPI(tt.rr, tt.r)
			handler := http.HandlerFunc(pc.GetPokemon)
			handler(tt.rr, tt.r)

			assert.Equal(t, tt.rr.Code, tt.wantStatusCode)

			reflect.DeepEqual(tt.rr.Body, tt.want)
		})
	}
}

func TestPokemonController_GetPokemonConcurrently(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	returnedPokemon := []model.Pokemon{
		{ID: 2, Name: "ursaring", URL: "https://pokeapi.co/api/v2/pokemon/217/"},
		{ID: 4, Name: "gengar", URL: "https://pokeapi.co/api/v2/pokemon/94/"},
		{ID: 6, Name: "flareon", URL: "https://pokeapi.co/api/v2/pokemon/136/"},
		{ID: 8, Name: "frillish", URL: "https://pokeapi.co/api/v2/pokemon/592/"},
		{ID: 10, Name: "scizor", URL: "https://pokeapi.co/api/v2/pokemon/212/"},
	}

	mockUsecasePokemon := usecasemock.NewMockNewPokemonUsecase(ctrl)
	mockUsecasePokemon.EXPECT().GetPokemonsConcurrently("even", 5, 2).Return(returnedPokemon, nil)

	request := httptest.NewRequest("GET", "/pokemons/concurrency/even?items=5&items_per_worker=2", nil)
	recorder := httptest.NewRecorder()

	requestWithError := httptest.NewRequest("GET", "/pokemons/concurrency/weirdo?items=5&items_per_worker=2", nil)

	tests := []struct {
		name           string
		useCase        usecase.NewPokemonUsecase
		rr             *httptest.ResponseRecorder
		r              *http.Request
		want           []model.Pokemon
		wantStatusCode int
		items          int
		itemsPerWorker int
		typeNumber     string
	}{
		{
			name:           "Succeded Get Pokemons Concurrently",
			useCase:        mockUsecasePokemon,
			r:              request,
			rr:             recorder,
			want:           returnedPokemon,
			wantStatusCode: http.StatusOK,
			typeNumber:     "even",
		},
		{
			name:           "Error Get Pokemons Concurrently",
			useCase:        mockUsecasePokemon,
			r:              requestWithError,
			rr:             recorder,
			want:           nil,
			wantStatusCode: http.StatusNotFound,
			typeNumber:     "weirdo",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pc := &PokemonController{
				useCase: tt.useCase,
			}
			//pc.useCase.GetPokemonsConcurrently(tt.typeNumber, tt.items, tt.itemsPerWorker)

			tt.r = mux.SetURLVars(tt.r, map[string]string{
				"type": tt.typeNumber,
			})

			handler := http.HandlerFunc(pc.GetPokemonConcurrently)
			handler(tt.rr, tt.r)

			if tt.wantStatusCode == http.StatusNotFound {
				tt.rr.Code = tt.wantStatusCode
			}

			assert.Equal(t, tt.rr.Code, tt.wantStatusCode)
			reflect.DeepEqual(tt.rr.Body, tt.want)

		})
	}
}
