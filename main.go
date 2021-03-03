package main

import (
	"github.com/dannegm/academy-go-q12021/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()

	app.GET("/pokedex", controller.GetPokemonList)

	app.Run()
}
