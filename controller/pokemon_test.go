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
)

func TestPokemonController_Index(t *testing.T) {

	request := httptest.NewRequest("GET", "/", nil)
	recorder := httptest.NewRecorder()

	tests := []struct {
		name string
		r    *http.Request
		rr   *httptest.ResponseRecorder
	}{
		{
			name: "Succeded Index Http Request",
			r:    request,
			rr:   recorder,
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
			if status := tt.rr.Code; status != http.StatusOK {
				t.Fatalf("handler returned wrong status code: got %v want %v",
					status, http.StatusOK)
			}

			// Check the response body is what we expect.
			expected := `{ "message": "Welcome to my Poke-API" }`
			if tt.rr.Body.String() != expected {
				t.Fatalf("handler returned unexpected body: got %v want %v",
					tt.rr.Body.String(), expected)
			}
		})
	}
}

func TestPokemonController_GetPokemons(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	pokemons := []model.Pokemon{
		{ID: 1, Name: "greninja", URL: "https://pokeapi.co/api/v2/pokemon/658/"},
		{ID: 2, Name: "ursaring", URL: "https://pokeapi.co/api/v2/pokemon/217/"},
		{ID: 3, Name: "arcanine", URL: "https://pokeapi.co/api/v2/pokemon/59/"},
		{ID: 4, Name: "gengar", URL: "https://pokeapi.co/api/v2/pokemon/94/"},
		{ID: 5, Name: "porygon", URL: "https://pokeapi.co/api/v2/pokemon/137/"},
	}

	mockUsecasePokemon := usecasemock.NewMockNewPokemonUsecase(ctrl)
	mockUsecasePokemon.EXPECT().GetPokemons().Return(pokemons, nil)

	request := httptest.NewRequest("GET", "/pokemons", nil)
	recorder := httptest.NewRecorder()

	tests := []struct {
		name    string
		useCase usecase.NewPokemonUsecase
		r       *http.Request
		rr      *httptest.ResponseRecorder
		want    []model.Pokemon
	}{
		{
			name:    "succeded Get Pokemons",
			useCase: mockUsecasePokemon,
			r:       request,
			rr:      recorder,
			want:    pokemons,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pc := &PokemonController{
				useCase: tt.useCase,
			}
			handler := http.HandlerFunc(pc.GetPokemons)
			handler.ServeHTTP(tt.rr, tt.r)

			if status := tt.rr.Code; status != http.StatusOK {
				t.Fatalf("handler returned wrong status code: got %v want %v",
					status, http.StatusOK)
			}

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
		name      string
		useCase   usecase.NewPokemonUsecase
		r         *http.Request
		rr        *httptest.ResponseRecorder
		want      model.Pokemon
		pokemonId int
	}{
		{
			name:      "Succeded Get Pokemon Id: 1",
			useCase:   mockUsecasePokemon,
			r:         request,
			rr:        recorder,
			want:      pokemon,
			pokemonId: 1,
		},
		{
			name:      "Succeded Get Pokemon Id: 3",
			useCase:   mockUsecasePokemon,
			r:         requestSecondTest,
			rr:        recorderSecondTest,
			want:      pokemonTest,
			pokemonId: 3,
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

			if status := tt.rr.Code; status != http.StatusOK {
				t.Fatalf("handler returned wrong status code: got %v want %v",
					status, http.StatusOK)
			}

			reflect.DeepEqual(tt.rr.Body, tt.want)
		})
	}
}

func TestPokemonController_GetPokemonsFromExternalAPI(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	newPokemons := &[]model.SinglePokeExternal{
		{Name: "delcatty", URL: "https://pokeapi.co/api/v2/pokemon/301/"},
		{Name: "sableye", URL: "https://pokeapi.co/api/v2/pokemon/302/"},
		{Name: "mawile", URL: "https://pokeapi.co/api/v2/pokemon/303/"},
		{Name: "aron", URL: "https://pokeapi.co/api/v2/pokemon/304/"},
		{Name: "lairon", URL: "https://pokeapi.co/api/v2/pokemon/305/"},
	}
	mockUsecasePokemon := usecasemock.NewMockNewPokemonUsecase(ctrl)
	mockUsecasePokemon.EXPECT().GetPokemonsFromExternalAPI().Return(newPokemons, nil)

	request := httptest.NewRequest("GET", "/pokemons/external", nil)
	recorder := httptest.NewRecorder()

	tests := []struct {
		name    string
		useCase usecase.NewPokemonUsecase
		r       *http.Request
		rr      *httptest.ResponseRecorder
		want    *[]model.SinglePokeExternal
	}{
		{
			name:    "Succeded Get Pokemon From External API",
			useCase: mockUsecasePokemon,
			r:       request,
			rr:      recorder,
			want:    newPokemons,
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

			if status := tt.rr.Code; status != http.StatusOK {
				t.Fatalf("handler returned wrong status code: got %v want %v",
					status, http.StatusOK)
			}

			reflect.DeepEqual(tt.rr.Body, tt.want)
		})
	}
}
