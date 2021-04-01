package interactor

import (
	"errors"
	"fmt"
	"github.com/ToteEmmanuel/academy-go-q12021/domain/model"
	"github.com/ToteEmmanuel/academy-go-q12021/usecase/presenter"
	"github.com/ToteEmmanuel/academy-go-q12021/usecase/repository"
	"github.com/go-resty/resty/v2"
	"sync"
)

type pokeInteractor struct {
	PokeRepo      repository.PokeRepository
	PokePresenter presenter.PokePresenter
	RestClient    *resty.Client
}

type PokeInteractor interface {
	Get(id int32) (*model.Pokemon, error)
	GetAll() ([]*model.Pokemon, error)
	CatchOne(int32) (*model.Pokemon, error)
	GetAllWorkers(string, int, int) ([]*model.Pokemon, error)
}

func NewPokeInteractor(r repository.PokeRepository, p presenter.PokePresenter, client *resty.Client) PokeInteractor {
	return &pokeInteractor{
		PokeRepo:      r,
		PokePresenter: p,
		RestClient:    client,
	}
}

func (pI *pokeInteractor) GetAllWorkers(typeQuery string, items, itemsPerWorker int) ([]*model.Pokemon, error) {
	p, err := pI.PokeRepo.FindAllWorkers(typeQuery, items, itemsPerWorker)
	if err != nil {
		return nil, err
	}
	return pI.PokePresenter.ResponsePokemons(p), nil
}

func (pI *pokeInteractor) Get(id int32) (*model.Pokemon, error) {
	p, err := pI.PokeRepo.FindById(id)
	if err != nil {
		return nil, err
	}
	return pI.PokePresenter.ResponsePokemon(p), nil
}

func (pI *pokeInteractor) GetAll() ([]*model.Pokemon, error) {
	p, err := pI.PokeRepo.FindAll()
	if err != nil {
		return nil, err
	}
	return pI.PokePresenter.ResponsePokemons(p), nil
}

func (pI *pokeInteractor) CatchOne(id int32) (*model.Pokemon, error) {
	p, err := pI.lookForPokemon(id)
	if err != nil {
		return nil, err
	}
	p, err = pI.PokeRepo.Save(p)
	if err != nil {
		return nil, err
	}
	return pI.PokePresenter.ResponsePokemon(p), nil
}

func (pI *pokeInteractor) lookForPokemon(id int32) (*model.Pokemon, error) {
	var waitgroup sync.WaitGroup
	waitgroup.Add(2)
	catchedPokemon := &model.Pokemon{
		Id:         id,
		Species:    "",
		Sprite:     "",
		FlavorText: "",
		Types:      nil,
	}
	var err error
	err = fetchBase(catchedPokemon, pI.RestClient)
	if err != nil {
		fmt.Println(err)
	}
	go func() {
		err = fetchSpecies(catchedPokemon, pI.RestClient)
		waitgroup.Done()
		fmt.Println("Fetched Species...")
	}()
	go func() {
		err = fetchBase(catchedPokemon, pI.RestClient)
		waitgroup.Done()
		fmt.Println("Fetched Base...")
	}()
	waitgroup.Wait()
	if err != nil {
		return nil, err
	}
	return catchedPokemon, nil
}

func fetchSpecies(pokemon *model.Pokemon, client *resty.Client) error {
	resp, err := client.R().
		SetPathParams(map[string]string{
			"pokeId": fmt.Sprint(pokemon.Id),
		}).
		SetResult(pokemonDefinitionDto{}).
		SetHeader("Accept", "application/json").
		Get("https://pokeapi.co/api/v2/pokemon-species/{pokeId}")
	if err != nil {
		return err
	}
	if resp.StatusCode() != 200 {
		return errors.New(
			fmt.Sprintf("Error in communication with downstream service. status:[%s]", resp.StatusCode()))
	}
	result := *resp.Result().(*pokemonDefinitionDto)
	pokemon.FlavorText = result.FlavorText[0]["flavor_text"].(string)
	return nil
}

func fetchBase(pokemon *model.Pokemon, client *resty.Client) error {
	resp, err := client.R().
		SetPathParams(map[string]string{
			"pokeId": fmt.Sprint(pokemon.Id),
		}).
		SetResult(pokemonBaseDto{}).
		SetHeader("Accept", "application/json").
		Get("https://pokeapi.co/api/v2/pokemon/{pokeId}/")
	if err != nil {
		return err
	}
	if resp.StatusCode() != 200 {
		return errors.New(
			fmt.Sprintf("Error in communication with downstream service. status:[%s]", resp.StatusCode()))
	}
	result := *resp.Result().(*pokemonBaseDto)
	types := []string{}
	for _, typeDto := range result.Types {
		types = append(types, typeDto.Type["name"])
	}
	pokemon.Sprite = result.Sprites["front_default"].(string)
	pokemon.Species = result.Species["name"]
	pokemon.Types = types
	return nil
}

type pokemonDefinitionDto struct {
	Id         int32                    `json:"id"`
	FlavorText []map[string]interface{} `json:"flavor_text_entries"`
}

type pokemonBaseDto struct {
	Id      int32                  `json:"id"`
	Sprites map[string]interface{} `json:"sprites"`
	Species map[string]string      `json:"species"`
	Types   []pokemonTypeDto       `json:"types"`
}

type pokemonTypeDto struct {
	Type map[string]string `json:"type"`
}
