package interactor

import (
	"errors"
	"fmt"
	"github.com/ToteEmmanuel/academy-go-q12021/domain/model"
	"github.com/go-resty/resty/v2"
)

//mockgen -source=usecase/interactor/infoclient.go  -destination=usecase/interactor/mock/infoclient_mock.go mock InfoClient
type InfoClient interface {
	FetchSpecies(pokemon *model.Pokemon, client *resty.Client) error
	FetchBase(pokemon *model.Pokemon, client *resty.Client) error
}

type infoClient struct{}

func NewInfoClient() InfoClient {
	return &infoClient{}
}

func (infoClient) FetchSpecies(pokemon *model.Pokemon, client *resty.Client) error {
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
			fmt.Sprintf("Error in communication with downstream service [species]. status:[%d]", resp.StatusCode()))
	}
	result := *resp.Result().(*pokemonDefinitionDto)
	pokemon.FlavorText = result.FlavorText[0]["flavor_text"].(string)
	return nil
}

func (infoClient) FetchBase(pokemon *model.Pokemon, client *resty.Client) error {
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
			fmt.Sprintf("Error in communication with downstream service [base]. status:[%d]", resp.StatusCode()))
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
	Id         int                      `json:"id"`
	FlavorText []map[string]interface{} `json:"flavor_text_entries"`
}

type pokemonBaseDto struct {
	Id      int                    `json:"id"`
	Sprites map[string]interface{} `json:"sprites"`
	Species map[string]string      `json:"species"`
	Types   []pokemonTypeDto       `json:"types"`
}

type pokemonTypeDto struct {
	Type map[string]string `json:"type"`
}
