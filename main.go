package main

import (
	"github.com/dannegm/academy-go-q12021/controller"
	"github.com/dannegm/academy-go-q12021/models"
	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()

	pokedex, _ := models.PokedexFromFile()

	app.GET("/pokedex", controller.GetPokemonList(pokedex))
	app.GET("/pokemon/:pokemonID", controller.GetPokemonByID(pokedex))

	app.Run()
}
