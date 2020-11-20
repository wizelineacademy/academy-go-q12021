package app

import (
	"fmt"
	"github.com/spf13/viper"
	"golang-bootcamp-2020/controllers"
	"golang-bootcamp-2020/infrastructure/router"
	"golang-bootcamp-2020/repository/db"
	"golang-bootcamp-2020/repository/rest"
	"golang-bootcamp-2020/services"
)

//Setting application
func StartApp() {
	viper.SetConfigName("config") // config file name without extension
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config/") // config file path
	viper.AutomaticEnv()             // read value ENV variable

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("fatal error config file: default \n", err)
		panic(err)
	}

	handler := controllers.NewAppController(services.NewService(rest.NewRickAndMortyApiRepository(), db.NewDbRepository()))

	if err := router.StartRouter(handler); err != nil {
		panic(err)
	}
}
