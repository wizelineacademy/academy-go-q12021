package datastore

import (
	"fmt"
	"github.com/ToteEmmanuel/academy-go-q12021/domain/model"
	"github.com/stretchr/testify/assert"
	"runtime"
	"testing"
)

func TestCsvPokeStorage_FindById(t *testing.T) {
	tables := []struct {
		scenario string
		storage  *CsvPokeStorage
		id       int
		expected *model.Pokemon
	}{
		{"Not found",
			&CsvPokeStorage{make(map[int]*model.Pokemon), ""},
			1,
			nil,
		},
		{"Found in one",
			&CsvPokeStorage{map[int]*model.Pokemon{
				1: {Id: 1},
			}, ""},
			1,
			&model.Pokemon{Id: 1},
		},
		{"Found in many",
			&CsvPokeStorage{map[int]*model.Pokemon{
				1: {Id: 1},
				2: {Id: 2},
				3: {Id: 3},
			}, ""},
			1,
			&model.Pokemon{Id: 1},
		},
	}
	for _, table := range tables {
		pokemon := table.storage.FindById(table.id)
		assert.Equal(t, pokemon, table.expected, fmt.Sprintf("Failed : %s", table.scenario))
	}
}

func TestCsvPokeStorage_FindAll(t *testing.T) {
	tables := []struct {
		scenario string
		storage  *CsvPokeStorage
		expected []*model.Pokemon
	}{
		{"Empty map",
			&CsvPokeStorage{make(map[int]*model.Pokemon), ""},
			[]*model.Pokemon{},
		},
		{"One in map",
			&CsvPokeStorage{map[int]*model.Pokemon{
				1: {Id: 1},
			}, ""},
			[]*model.Pokemon{&model.Pokemon{Id: 1}},
		},
		{"Many in map",
			&CsvPokeStorage{map[int]*model.Pokemon{
				1: {Id: 1},
				2: {Id: 2},
				3: {Id: 3},
			}, ""},
			[]*model.Pokemon{
				&model.Pokemon{Id: 1},
				&model.Pokemon{Id: 2},
				&model.Pokemon{Id: 3},
			},
		},
	}
	for _, table := range tables {
		pokemons := table.storage.FindAll()
		assert.Equal(t, len(pokemons), len(table.expected), fmt.Sprintf("Failed : %s", table.scenario))
	}
}

func TestCsvPokeStorage_FindAllWorkers(t *testing.T) {

	fourMaxPoolSizeMap := make(map[int]*model.Pokemon)
	poolSize := runtime.GOMAXPROCS(0)

	for i := 1; i <= poolSize*4; i++ {
		fourMaxPoolSizeMap[i] = &model.Pokemon{Id: i}
	}

	tables := []struct {
		scenario       string
		storage        *CsvPokeStorage
		typeStr        string
		items          int
		itemsPerWorker int
		amount         int
	}{
		{"Empty map , odd, 1 items, 1 itemsperworker",
			&CsvPokeStorage{make(map[int]*model.Pokemon), ""},
			"odd",
			1,
			1,
			0,
		},
		{"none matching in map , odd, 1 items, 1 itemsperworker",
			&CsvPokeStorage{map[int]*model.Pokemon{
				2: {Id: 2},
				4: {Id: 4},
				6: {Id: 6},
			}, ""},
			"odd",
			2,
			2,
			0,
		},
		{"many matching in map , odd, 2 items, 2 itemsperworker",
			&CsvPokeStorage{fourMaxPoolSizeMap, ""},
			"odd",
			2,
			2,
			2,
		},
		{"many matching in map , even, 2 items, 2 itemsperworker",
			&CsvPokeStorage{fourMaxPoolSizeMap, ""},
			"even",
			2,
			2,
			2,
		},
		{"requested more than workers can get , even, poolsize# items, 1 itemsperworker",
			&CsvPokeStorage{fourMaxPoolSizeMap, ""},
			"even",
			2 * poolSize,
			1,
			poolSize,
		},
	}
	for _, table := range tables {
		pokemons, _ := table.storage.FindAllWorkers(table.typeStr, table.items, table.itemsPerWorker)
		assert.Equal(t, table.amount, len(pokemons), fmt.Sprintf("Failed : %s", table.scenario))
		if table.amount > 0 {
			for _, pokemon := range pokemons {
				if table.typeStr == "odd" {
					assert.True(t, pokemon.Id%2 != 0, fmt.Sprintf("Failed type in : %s", table.scenario))
				} else {
					assert.True(t, pokemon.Id%2 == 0, fmt.Sprintf("Failed type in : %s", table.scenario))
				}
			}
		}
	}
}
