package controller

import (
	"errors"
	"fmt"
	"github.com/ToteEmmanuel/academy-go-q12021/domain/model"
	mock_interactor "github.com/ToteEmmanuel/academy-go-q12021/usecase/interactor/mock"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_pokeController_CatchPokemon(t *testing.T) {
	tables := []struct {
		scenario string
		id       string
		expected *model.Pokemon
		error    error
		status   int
	}{
		{"Error in interactor",
			"1",
			nil,
			errors.New("not found"),
			http.StatusInternalServerError,
		},
		{"Found",
			"1",
			&model.Pokemon{Id: 1},
			nil,
			http.StatusOK,
		},
	}
	for _, table := range tables {
		interactor := getMockPokeInteractor(t)
		interactor.EXPECT().CatchOne(gomock.Any()).Return(table.expected, table.error).Times(1)
		c := NewPokeController(interactor)
		req, err := http.NewRequest("GET",
			fmt.Sprintf("/pokemon/%s/catch", table.id),
			nil)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(c.CatchPokemon)
		handler.ServeHTTP(rr, req)
		assert.Equal(t, table.status, rr.Code, "Wrong Response Code, %s", table.scenario)
		contentType := rr.Result().Header.Get("Content-Type")
		assert.Equal(t, "Application/Json", contentType, "Wrong content type %s", table.scenario)
	}
}

func Test_pokeController_GetPokemon(t *testing.T) {
	tables := []struct {
		scenario string
		id       string
		expected *model.Pokemon
		error    error
		status   int
	}{
		{"Not found",
			"1",
			nil,
			errors.New("not found"),
			http.StatusNotFound,
		},
		{"Found",
			"1",
			&model.Pokemon{Id: 1},
			nil,
			http.StatusOK,
		},
	}
	for _, table := range tables {
		interactor := getMockPokeInteractor(t)
		interactor.EXPECT().Get(gomock.Any()).Return(table.expected, table.error).Times(1)
		c := NewPokeController(interactor)
		req, err := http.NewRequest("GET",
			fmt.Sprintf("/pokemon/%s", table.id),
			nil)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(c.GetPokemon)
		handler.ServeHTTP(rr, req)
		assert.Equal(t, table.status, rr.Code, "Wrong Response Code, %s", table.scenario)
		contentType := rr.Result().Header.Get("Content-Type")
		assert.Equal(t, "Application/Json", contentType, "Wrong content type %s", table.scenario)
	}
}

func Test_pokeController_GetPokemons(t *testing.T) {
	tables := []struct {
		scenario string
		expected []*model.Pokemon
		error    error
		status   int
	}{
		{"Empty",
			[]*model.Pokemon{},
			nil,
			http.StatusOK,
		},
		{"Values",
			[]*model.Pokemon{
				&model.Pokemon{Id: 1},
				&model.Pokemon{Id: 3},
				&model.Pokemon{Id: 5},
			},
			nil,
			http.StatusOK,
		},
		{"Error Processing",
			nil,
			errors.New("Random error."),
			http.StatusInternalServerError,
		},
	}
	for _, table := range tables {
		interactor := getMockPokeInteractor(t)
		interactor.EXPECT().GetAll().Return(table.expected, table.error)
		c := NewPokeController(interactor)
		req, err := http.NewRequest("GET", "/pokemon", nil)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(c.GetPokemons)
		handler.ServeHTTP(rr, req)
		assert.Equal(t, table.status, rr.Code, "Wrong Response Code, %s", table.scenario)
		contentType := rr.Result().Header.Get("Content-Type")
		assert.Equal(t, "Application/Json", contentType, "Wrong content type %s", table.scenario)
	}
}

func TestPokeController_GetPokemonsWithWorkers(t *testing.T) {
	tables := []struct {
		scenario       string
		expected       []*model.Pokemon
		error          error
		status         int
		typeStr        string
		items          string
		itemsPerWorker string
	}{
		{"Empty",
			[]*model.Pokemon{},
			nil,
			http.StatusOK,
			"odd",
			"1",
			"1",
		},
		{"Values",
			[]*model.Pokemon{
				&model.Pokemon{Id: 1},
				&model.Pokemon{Id: 3},
				&model.Pokemon{Id: 5},
			},
			nil,
			http.StatusOK,
			"odd",
			"1",
			"1",
		},
		{"Error Processing",
			nil,
			errors.New("Random error."),
			http.StatusInternalServerError,
			"odd",
			"1",
			"1",
		},
		{"malformed type",
			nil,
			errors.New("Random error."),
			http.StatusBadRequest,
			"oddo",
			"1",
			"1",
		},
		{"Wrong items type",
			nil,
			errors.New("Random error."),
			http.StatusBadRequest,
			"odd",
			"nope",
			"1",
		},
		{"wrong itemsPerWorker",
			nil,
			errors.New("Random error."),
			http.StatusBadRequest,
			"odd",
			"1",
			"jojo",
		},
	}
	for _, table := range tables {
		interactor := getMockPokeInteractor(t)
		interactor.EXPECT().GetAllWorkers(gomock.Any(), gomock.Any(), gomock.Any()).Return(table.expected, table.error)
		c := NewPokeController(interactor)
		req, err := http.NewRequest("GET",
			fmt.Sprintf("/pokemon/workers"),
			nil)
		req = mux.SetURLVars(req, map[string]string{
			"type":              table.typeStr,
			"items":             table.items,
			"items_per_workers": table.itemsPerWorker,
		})

		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(c.GetPokemonsWithWorkers)
		handler.ServeHTTP(rr, req)
		assert.Equal(t, table.status, rr.Code, "Wrong Response Code, %s", table.scenario)
		contentType := rr.Result().Header.Get("Content-Type")
		assert.Equal(t, "Application/Json", contentType, "Wrong content type %s", table.scenario)
	}
}

func getMockPokeInteractor(t *testing.T) *mock_interactor.MockPokeInteractor {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	return mock_interactor.NewMockPokeInteractor(ctrl)
}
