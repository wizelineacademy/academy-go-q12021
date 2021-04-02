package repository

import (
	"errors"
	"fmt"
	"github.com/ToteEmmanuel/academy-go-q12021/domain/model"
	"github.com/ToteEmmanuel/academy-go-q12021/infrastructure/datastore/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPokeRepository_FindById(t *testing.T) {
	storage := getMockPokeStorage(t)
	pr := NewPokeRepository(storage)
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
		storage.EXPECT().FindById(gomock.Eq(table.id)).Return(table.expected)
		pokemon, err := pr.FindById(table.id)
		assert.Equal(t, pokemon, table.expected, fmt.Sprintf("Failed : %s", table.scenario))
		assert.Equal(t, err, table.error, fmt.Sprintf("Failed : %s", table.scenario))
	}
}

func TestPokeRepository_FindAll(t *testing.T) {
	storage := getMockPokeStorage(t)
	pr := NewPokeRepository(storage)
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
				&model.Pokemon{Id: 2},
				&model.Pokemon{Id: 3},
			},
			nil,
		},
	}
	for _, table := range tables {
		storage.EXPECT().FindAll().Return(table.expected)
		pokemon, err := pr.FindAll()
		assert.Equal(t, len(pokemon), len(table.expected), fmt.Sprintf("Failed : %s", table.scenario))
		assert.Equal(t, err, table.error, fmt.Sprintf("Failed : %s", table.scenario))
	}
}

func TestPokeRepository_FindAllWorkers(t *testing.T) {
	storage := getMockPokeStorage(t)
	pr := NewPokeRepository(storage)
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
		storage.EXPECT().FindAllWorkers(gomock.Eq("odd"), gomock.Eq(1), gomock.Eq(1)).
			Return(table.expected, table.error)
		pokemon, err := pr.FindAllWorkers("odd", 1, 1)
		assert.Equal(t, len(pokemon), len(table.expected), fmt.Sprintf("Failed : %s", table.scenario))
		assert.Equal(t, err, table.error, fmt.Sprintf("Failed : %s", table.scenario))
	}
}

func TestPokeRepository_Save(t *testing.T) {
	storage := getMockPokeStorage(t)
	pr := NewPokeRepository(storage)
	tables := []struct {
		scenario string
		id       int
		expected *model.Pokemon
		error    error
	}{
		{"Save",
			1,
			&model.Pokemon{Id: 1},
			nil,
		},
		{"Error Saving",
			1,
			nil,
			errors.New("not found"),
		},
	}
	for _, table := range tables {
		storage.EXPECT().Save(gomock.Any()).Return(table.expected, table.error)
		pokemon, err := pr.Save(table.expected)

		if pokemon != nil {
			assert.Equal(t, table.expected.Id, pokemon.Id, fmt.Sprintf("Failed : %s", table.scenario))
		} else {
			assert.Equal(t, table.expected, pokemon, fmt.Sprintf("Failed : %s", table.scenario))
		}
		assert.Equal(t, table.error, err, fmt.Sprintf("Failed : %s", table.scenario))
	}
}

func getMockPokeStorage(t *testing.T) *mock.MockPokeStorage {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	return mock.NewMockPokeStorage(ctrl)
}
