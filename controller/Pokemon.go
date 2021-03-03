package controller

import (
	"strconv"

	"github.com/dannegm/academy-go-q12021/models"
	"github.com/gin-gonic/gin"
)

// PokedexList stores a list of Pokemon
type PokedexList []models.Pokemon

// PokemonListResponse for the response
type PokemonListResponse struct {
	Pokedex PokedexList `json:"data"`
}

// PokemonResponse for the response
type PokemonResponse struct {
	Pokemon models.Pokemon `json:"data"`
}

// ErrorResponse for error handling
type ErrorResponse struct {
	Status  int    `json:"satus"`
	Message string `json:"message"`
}

// GetPokemonList Pokes router
func GetPokemonList(pokedex models.Pokedex) func(*gin.Context) {
	return func(context *gin.Context) {
		pokedexList := PokedexList{}
		for _, pokemon := range pokedex {
			pokedexList = append(pokedexList, pokemon)
		}

		context.JSON(200, PokemonListResponse{
			Pokedex: pokedexList,
		})
	}
}

// GetPokemonByID to get a single Pokemon filter by ID
func GetPokemonByID(pokedex models.Pokedex) func(*gin.Context) {
	return func(context *gin.Context) {
		pokemonID, err := strconv.Atoi(context.Param("pokemonID"))

		if err != nil {
			context.JSON(400, ErrorResponse{
				Status:  400,
				Message: "Invalid Pokemon ID",
			})
		} else {
			pokemon := pokedex[pokemonID]

			context.JSON(200, PokemonResponse{
				Pokemon: pokemon,
			})
		}

	}
}
