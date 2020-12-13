package app

import (
	"fmt"

	"golang-bootcamp-2020/controllers"
	"golang-bootcamp-2020/infrastructure/router"
	"golang-bootcamp-2020/repository/db"
	"golang-bootcamp-2020/repository/rest"
	"golang-bootcamp-2020/services"

	"github.com/go-resty/resty/v2"
	"github.com/spf13/viper"
)

//StartApp - Setting application
func StartApp() {
	viper.SetConfigName("config") // config file name without extension
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./src/config/") // config file path
	viper.AutomaticEnv()                 // read value ENV variable

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("fatal error config file: default \n", err)
		panic(err)
	}

	restClient := resty.New()
	//Checking health of API
	if response, err := restClient.R().Get(viper.GetString("rest.host")); err != nil || response.StatusCode() != 200 {
		panic("external rest API not available")
	}

	handler := controllers.NewAppController(services.NewService(rest.NewRickAndMortyAPIRepository(restClient), db.NewDbRepository()))

	if err := router.StartRouter(handler); err != nil {
		panic(err)
	}
}
