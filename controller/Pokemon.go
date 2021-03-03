package controller

import (
	"github.com/dannegm/academy-go-q12021/models"
	"github.com/gin-gonic/gin"
)

type pokemonResponse struct {
	Pokedex models.Pokedex `json:"data"`
}

// GetPokemonList Pokes router
func GetPokemonList(context *gin.Context) {
	pokes, err := models.PokedexFromFile()

	if err != nil {
		context.Status(500)
	}

	response := pokemonResponse{
		Pokedex: pokes,
	}

	context.JSON(200, response)
}
