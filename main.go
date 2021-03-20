package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/halarcon-wizeline/academy-go-q12021/domain"
	"github.com/halarcon-wizeline/academy-go-q12021/usecase"
	"github.com/halarcon-wizeline/academy-go-q12021/controller"
	"github.com/halarcon-wizeline/academy-go-q12021/router"

	"github.com/sirupsen/logrus"
	"github.com/unrolled/render"
)

func main() {

	// logger setup
	logger, err := createLogger()
	if err != nil || logger == nil {
		log.Fatal("creating logger: %w", err)
	}

	// Create client
	pokemon := domain.NewPokemon(1, "Mewto")
	// fmt.Println(pokemon.Name)

	// Usecase
	useCase := usecase.New(*pokemon)

	// Controllers
	controller := controller.New(useCase, logger, render.New())

	// Setup application routes
	httpRouter := router.New(controller)

	var serverPort string = "8080"
	fmt.Println("Server listen at http://localhost" + ":" + serverPort)
	log.Fatal(http.ListenAndServe("localhost:"+serverPort, httpRouter))
}

func createLogger() (*logrus.Logger, error) {
	logLevel := "DEBUG"
	level, err := logrus.ParseLevel(logLevel)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"log_level": logLevel,
		}).Error("parsing log_level")

		return nil, err
	}

	logger := logrus.New()
	logger.SetLevel(level)
	logger.Out = os.Stdout
	logger.Formatter = &logrus.JSONFormatter{}
	return logger, nil
}
