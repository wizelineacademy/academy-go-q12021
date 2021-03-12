package presenter

import (
	"fmt"

	"github.com/AlejandroSeguraWIZ/academy-go-q12021/domain/model"
)

type pokemonPresenter struct{}

type PokemonPresenter interface {
	ResponsePokemons([]model.Pokemon) []model.PokemonResponse
}

func NewPokemonPresenter() PokemonPresenter {
	return &pokemonPresenter{}
}

func (p *pokemonPresenter) ResponsePokemons(pokemons []model.Pokemon) []model.PokemonResponse {
	result := []model.PokemonResponse{}
	for _, p := range pokemons {
		pf := model.PokemonResponse{
			Name:       p.Name,
			Cathegory:  fmt.Sprintf("%s/%s", p.Type, p.Subtype),
			Generation: p.Generation,
			Legendary:  p.Legendary,
			Score: model.PokemonScore{
				HP:      p.HP,
				Attack:  p.Attack,
				Defense: p.Defense,
				SpAtk:   p.SpAtk,
				SpDef:   p.SpDef,
				Speed:   p.Speed,
				Total:   p.Total,
			},
		}
		result = append(result, pf)
	}
	return result
}
