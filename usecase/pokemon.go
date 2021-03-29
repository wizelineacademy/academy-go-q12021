package usecase

// go:generate mockgen -source=usecase/pokemon.go -destination=usecase/mock/pokemon_usecase.go -package=mock

import (
	"math"
	"sync"

	"pokeapi/model"
	csvservice "pokeapi/service/csv"
	httpservice "pokeapi/service/http"
)

type PokemonUsecase struct {
	csvService  csvservice.NewCsvService
	httpService httpservice.NewHttpService
}

type NewPokemonUsecase interface {
	GetPokemons() ([]model.Pokemon, *model.Error)
	GetPokemonsConcurrently(typeNumber string, items int, itemsPerWorker int) ([]model.Pokemon, *model.Error)
	GetPokemon(pokemonId int) (model.Pokemon, *model.Error)
	GetPokemonsFromExternalAPI() (*[]model.SinglePokeExternal, *model.Error)
}

func New(s csvservice.NewCsvService, h httpservice.NewHttpService) *PokemonUsecase {
	return &PokemonUsecase{s, h}
}

func (us *PokemonUsecase) GetPokemons() ([]model.Pokemon, *model.Error) {
	return us.csvService.GetPokemons()
}

func (us *PokemonUsecase) GetPokemon(pokemonId int) (model.Pokemon, *model.Error) {
	return us.csvService.GetPokemon(pokemonId)
}

func calculatePoolSize(items int, itemsPerWorker int, totalPokemons int) int {
	var poolSize int
	if items%itemsPerWorker != 0 {
		poolSize = int(math.Ceil(float64(items) / float64(itemsPerWorker)))
	} else {
		poolSize = int(items / itemsPerWorker)
	}

	// If we overpass the number of workers above the half of number
	// of items it's gonna get into an infinit looop
	if poolSize > (totalPokemons / 2) {
		poolSize = totalPokemons / 2
	}
	return poolSize
}

func calculateMaxPokemons(totalPokemons int) int {
	var maxPokemons int

	if totalPokemons%2 == 0 {
		maxPokemons = totalPokemons / 2
	} else {
		maxPokemons = totalPokemons/2 + 1
	}
	return maxPokemons
}

func (us *PokemonUsecase) GetPokemonsConcurrently(typeNumber string, items int, itemsPerWorker int) ([]model.Pokemon, *model.Error) {
	pokemons, err := us.csvService.GetPokemons()

	if err != nil {
		return nil, err
	}

	filterType := model.TypeNumberFilter{
		Even: "even",
		Odd:  "odd",
	}
	totalPokemons := len(pokemons)
	poolSize := calculatePoolSize(items, itemsPerWorker, totalPokemons)
	maxPokemons := calculateMaxPokemons(totalPokemons)

	values := make(chan int)
	jobs := make(chan int, poolSize)
	shutdown := make(chan struct{})

	startIndex := 0
	var limit int
	limit = int(math.Ceil(float64(totalPokemons) / float64(poolSize)))
	lastLimit := (totalPokemons % limit)

	var wg sync.WaitGroup
	wg.Add(poolSize)

	for i := 0; i < poolSize; i++ {
		go func(jobs <-chan int) {
			for {
				var id int
				var limitRecalculated int
				start := <-jobs

				// We do need to iterate with the same limit every time.
				// on the last cycle we use the leftovers of the division (modulus)
				if limit+start >= totalPokemons && lastLimit != 0 { // lastLimit can be 0, take care of that
					limitRecalculated = start + lastLimit
				} else {
					limitRecalculated = start + limit
				}

				for j := start; j < limitRecalculated; j++ {
					id = pokemons[j].ID

					select {
					case values <- id:
					case <-shutdown:
						wg.Done()
						return
					}
				}
			}
		}(jobs)
	}

	for i := 0; i < poolSize; i++ {
		jobs <- startIndex
		startIndex += limit
	}
	close(jobs)

	var filteredPokemons []model.Pokemon = nil
	bucket := make(map[int]int, totalPokemons+1)
	for elem := range values {
		if typeNumber == filterType.Odd {
			if elem%2 != 0 && bucket[elem] == 0 {
				filteredPokemons = append(filteredPokemons, pokemons[elem-1])
				bucket[elem] = elem // we use the map to mark the ones that has been added to the collection
			}
		} else if typeNumber == filterType.Even {
			if elem%2 == 0 && bucket[elem] == 0 {
				filteredPokemons = append(filteredPokemons, pokemons[elem-1])
				bucket[elem] = elem // we use the map to mark the ones that has been added to the collection
			}
		}
		if len(filteredPokemons) >= items || len(filteredPokemons) >= maxPokemons {
			break // Finally if we reahc the items value or the possibly half that we cna take, break the loop
		}
	}

	// closing this channel we send the signal to all the goroutines to be finished
	close(shutdown)

	wg.Wait()

	return filteredPokemons, nil
}

func (us *PokemonUsecase) GetPokemonsFromExternalAPI() (*[]model.SinglePokeExternal, *model.Error) {
	newPokemons, err := us.httpService.GetPokemons()

	if err != nil {
		return nil, err
	}

	errorCsv := us.csvService.SavePokemons(&newPokemons)

	if errorCsv != nil {
		return nil, errorCsv
	}

	return &newPokemons, nil
}
