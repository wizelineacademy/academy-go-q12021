package interactor

import (
	"errors"
	"fmt"
	"github.com/ToteEmmanuel/academy-go-q12021/domain/model"
	mock_interactor "github.com/ToteEmmanuel/academy-go-q12021/usecase/interactor/mock"
	mock_repo "github.com/ToteEmmanuel/academy-go-q12021/usecase/repository/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"log"
	"strings"
	"testing"
)

func TestPokeInteractor_Get(t *testing.T) {
	repository := getMockPokeRepository(t)
	infoClient := getMockInfoClient(t)
	pI := NewPokeInteractor(repository, nil, infoClient)
	tables := []struct {
		scenario string
		id       int
		expected *model.Pokemon
		error    error
	}{
		{"Not found",
			1,
			nil,
			errors.New("not found"),
		},
		{"Found",
			1,
			&model.Pokemon{Id: 1},
			nil,
		},
	}
	for _, table := range tables {
		repository.EXPECT().FindById(gomock.Eq(table.id)).Return(table.expected, table.error)
		pokemon, err := pI.Get(table.id)
		assert.Equal(t, pokemon, table.expected, fmt.Sprintf("Failed : %s", table.scenario))
		assert.Equal(t, err, table.error, fmt.Sprintf("Failed : %s", table.scenario))
	}
}

func TestPokeInteractor_CatchOne(t *testing.T) {
	tables := []struct {
		scenario     string
		id           int
		expected     *model.Pokemon
		errorBase    error
		errorSpecies error
		errorRepo    error
	}{
		{"Not found",
			1,
			nil,
			errors.New("base error"),
			errors.New("fetch error"),
			nil,
		},
		{"Found",
			1,
			&model.Pokemon{Id: 1},
			nil,
			nil,
			nil,
		},
		{"Save Error",
			1,
			nil,
			nil,
			nil,
			errors.New("save error"),
		},
	}
	for _, table := range tables {
		//for some reason the Save mock was not working ok until I changed this init code here.
		repository := getMockPokeRepository(t)
		infoClient := getMockInfoClient(t)
		pI := NewPokeInteractor(repository, nil, infoClient)
		infoClient.EXPECT().FetchBase(gomock.Any(), gomock.Any()).Return(table.errorBase).Times(1)
		infoClient.EXPECT().FetchSpecies(gomock.Any(), gomock.Any()).Return(table.errorSpecies).Times(1)
		repository.EXPECT().Save(gomock.Any()).Return(table.expected, table.errorRepo)
		pokemon, err := pI.CatchOne(table.id)
		if pokemon != nil {
			assert.Equal(t, table.expected.Id, pokemon.Id, fmt.Sprintf("Failed : %s", table.scenario))
		} else {
			assert.Equal(t, table.expected, pokemon, fmt.Sprintf("Failed : %s", table.scenario))
		}
		log.Println(err)
		if table.errorBase != nil {
			assert.True(t, strings.Contains(fmt.Sprint(err), fmt.Sprint(table.errorBase)), fmt.Sprintf("Failed : %s", table.scenario))
		}
		if table.errorSpecies != nil {
			assert.True(t, strings.Contains(fmt.Sprint(err), fmt.Sprint(table.errorSpecies)), fmt.Sprintf("Failed : %s", table.scenario))
		}
		if table.errorRepo != nil {
			assert.True(t, strings.Contains(fmt.Sprint(err), fmt.Sprint(table.errorRepo)), fmt.Sprintf("Failed : %s", table.scenario))
		}
	}
}

func TestPokeInteractor_GetAll(t *testing.T) {
	repository := getMockPokeRepository(t)
	infoClient := getMockInfoClient(t)
	pI := NewPokeInteractor(repository, nil, infoClient)
	tables := []struct {
		scenario string
		expected []*model.Pokemon
		error    error
	}{
		{"Empty",
			[]*model.Pokemon{},
			nil,
		},
		{"Values",
			[]*model.Pokemon{
				&model.Pokemon{Id: 1},
				&model.Pokemon{Id: 3},
				&model.Pokemon{Id: 5},
			},
			nil,
		},
		{"Error Processing",
			nil,
			errors.New("Random error."),
		},
	}
	for _, table := range tables {
		repository.EXPECT().FindAll().
			Return(table.expected, table.error)
		pokemon, err := pI.GetAll()
		assert.Equal(t, len(pokemon), len(table.expected), fmt.Sprintf("Failed : %s", table.scenario))
		assert.Equal(t, err, table.error, fmt.Sprintf("Failed : %s", table.scenario))
	}
}

func TestPokeInteractor_GetAllWorkers(t *testing.T) {
	repository := getMockPokeRepository(t)
	infoClient := getMockInfoClient(t)
	pI := NewPokeInteractor(repository, nil, infoClient)
	tables := []struct {
		scenario string
		expected []*model.Pokemon
		error    error
	}{
		{"Empty",
			[]*model.Pokemon{},
			nil,
		},
		{"Values",
			[]*model.Pokemon{
				&model.Pokemon{Id: 1},
				&model.Pokemon{Id: 3},
				&model.Pokemon{Id: 5},
			},
			nil,
		},
		{"Error Processing",
			nil,
			errors.New("Random error."),
		},
	}
	for _, table := range tables {
		repository.EXPECT().FindAllWorkers(gomock.Eq("odd"), gomock.Eq(1), gomock.Eq(1)).
			Return(table.expected, table.error)
		pokemon, err := pI.GetAllWorkers("odd", 1, 1)
		assert.Equal(t, len(pokemon), len(table.expected), fmt.Sprintf("Failed : %s", table.scenario))
		assert.Equal(t, err, table.error, fmt.Sprintf("Failed : %s", table.scenario))
	}
}

func getMockPokeRepository(t *testing.T) *mock_repo.MockPokeRepository {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	return mock_repo.NewMockPokeRepository(ctrl)
}

func getMockInfoClient(t *testing.T) *mock_interactor.MockInfoClient {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	return mock_interactor.NewMockInfoClient(ctrl)
}
