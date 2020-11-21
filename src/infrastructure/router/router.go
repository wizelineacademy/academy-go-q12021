package router

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"golang-bootcamp-2020/controllers"
)

var (
	router = gin.Default()
)

//Setting api router
func StartRouter(c controllers.AppController) error {

	router.GET("health", c.GetHealth)

	router.GET("data/fetch", c.FetchData)

	router.GET("api/character/:id", c.GetCharacterById)
	router.GET("api/characters", c.GetCharacters)
	router.GET("api/findId/:name", c.GetCharacterIdByName)

	return router.Run(viper.GetString("app.port"))
}
