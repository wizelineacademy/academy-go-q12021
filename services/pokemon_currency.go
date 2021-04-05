package services

import (
	"sync"

	"github.com/cesararredondow/academy-go-q12021/models"
)

var (
	rejected = 0
	accepted = 0
)

// GetRegistries get the registries by the requested parameters
func (s *Service) GetRegistries(odd bool, itemsNumber int, workers int, pokemons []*models.Pokemon) ([]*models.Pokemon, error) {

	numberOfWorkers := itemsNumber / workers

	if itemsNumber%2 != 0 {
		numberOfWorkers++
	}

	jobs := make(chan models.Pokemon, len(pokemons))
	results := make(chan models.Pokemon, itemsNumber)
	registries := []*models.Pokemon{}

	wg := new(sync.WaitGroup)

	for _, pokemon := range pokemons {
		jobs <- *pokemon
	}
	close(jobs)

	wg.Add(numberOfWorkers)

	for i := 0; i < numberOfWorkers; i++ {
		go s.getPokemonsByRules(jobs, results, wg, odd, len(pokemons), i)
	}
	go func() {
		wg.Wait()
		close(results)
	}()

	//Add up the results from the results channel.
	for v := range results {
		p := new(models.Pokemon)
		p.ID = v.ID
		p.Name = v.Name
		registries = append(registries, p)
	}

	return registries, nil
}

//getPokemonsByRules evaluate the rules in a currency function
func (s *Service) getPokemonsByRules(jobs <-chan models.Pokemon, results chan<- models.Pokemon, wg *sync.WaitGroup, odd bool, pokemonsQuantity int, chanelID int) {
	// Decreasing internal counter for wait-group as soon as goroutine finishes
	defer wg.Done()
	for col := range jobs {
		if cap(results) == accepted {
			break
		}
		if pokemonsQuantity == (rejected + accepted) {
			break
		}
		if odd && col.ID%2 == 0 {
			results <- col
			accepted++
		}

		if !odd && col.ID%2 != 0 {
			results <- col
			accepted++
		}

		rejected++
	}
}
