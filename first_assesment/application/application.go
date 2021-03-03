package application

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/wizelineacademy/academy-go-q12021/router"

	"github.com/spf13/viper"
)

// RunApplication boot application that contains configuration and inject routes
func RunApplication() {
	r := router.NewRouting()

	viper.SetDefault("Address", "localhost:8080")
	viper.SetDefault("WriteTimeout", "15")
	viper.SetDefault("ReadTimeout", "15")
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	err := viper.ReadInConfig()

	if err != nil {
		panic(err)
	}
	writeTimeout, err := strconv.Atoi(viper.Get("WriteTimeout").(string))
	if err != nil {
		panic(err)
	}
	readTimeout, err := strconv.Atoi(viper.Get("ReadTimeout").(string))
	if err != nil {
		panic(err)
	}
	server := &http.Server{
		Handler:      r,
		Addr:         viper.Get("Address").(string),
		WriteTimeout: time.Duration(writeTimeout) * time.Second,
		ReadTimeout:  time.Duration(readTimeout) * time.Second,
	}
	log.Fatal(server.ListenAndServe())
}
