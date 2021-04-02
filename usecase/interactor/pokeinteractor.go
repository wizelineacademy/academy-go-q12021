package interactor

import (
	"fmt"
	"github.com/ToteEmmanuel/academy-go-q12021/domain/model"
	"github.com/ToteEmmanuel/academy-go-q12021/usecase/repository"
	"github.com/go-resty/resty/v2"
	"sync"
)

type pokeInteractor struct {
	PokeRepo   repository.PokeRepository
	RestClient *resty.Client
	InfoClient InfoClient
}

//mockgen -source=usecase/interactor/pokeinteractor.go  -destination=usecase/interactor/mock/pokeinteractor_mock.go mock PokeInteractor
type PokeInteractor interface {
	Get(id int) (*model.Pokemon, error)
	GetAll() ([]*model.Pokemon, error)
	CatchOne(int) (*model.Pokemon, error)
	GetAllWorkers(string, int, int) ([]*model.Pokemon, error)
}

func NewPokeInteractor(r repository.PokeRepository, client *resty.Client, infoClient InfoClient) PokeInteractor {
	return &pokeInteractor{
		PokeRepo:   r,
		RestClient: client,
		InfoClient: infoClient,
	}
}

func (pI *pokeInteractor) GetAllWorkers(typeQuery string, items, itemsPerWorker int) ([]*model.Pokemon, error) {
	p, err := pI.PokeRepo.FindAllWorkers(typeQuery, items, itemsPerWorker)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (pI *pokeInteractor) Get(id int) (*model.Pokemon, error) {
	p, err := pI.PokeRepo.FindById(id)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (pI *pokeInteractor) GetAll() ([]*model.Pokemon, error) {
	p, err := pI.PokeRepo.FindAll()
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (pI *pokeInteractor) CatchOne(id int) (*model.Pokemon, error) {
	p, err := pI.lookForPokemon(id)
	if err != nil {
		return nil, err
	}
	p, err = pI.PokeRepo.Save(p)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (pI *pokeInteractor) lookForPokemon(id int) (*model.Pokemon, error) {
	var waitgroup sync.WaitGroup
	waitgroup.Add(2)
	caughtPokemon := &model.Pokemon{
		Id:         id,
		Species:    "",
		Sprite:     "",
		FlavorText: "",
		Types:      nil,
	}
	errChan := make(chan error)
	defer close(errChan)
	go func() {
		err := pI.InfoClient.FetchSpecies(caughtPokemon, pI.RestClient)
		waitgroup.Done()
		errChan <- err
		fmt.Println("Fetched Species...")
	}()
	go func() {
		err := pI.InfoClient.FetchBase(caughtPokemon, pI.RestClient)
		waitgroup.Done()
		errChan <- err
		fmt.Println("Fetched Base...")
	}()

	waitgroup.Wait()
	var commError error
	for i := 0; i < 2; i++ {
		err := <-errChan
		if err != nil {
			if commError != nil {
				commError = fmt.Errorf("%v%w\n", commError, err)
				continue
			}
			commError = fmt.Errorf("communication error while fetching \n%w\n", err)
		}
	}
	if commError != nil {
		return nil, commError
	}
	return caughtPokemon, nil
}
