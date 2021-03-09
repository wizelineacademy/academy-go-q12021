package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gorilla/mux"

	"github.com/jesus-mata/academy-go-q12021/infrastructure"
	"github.com/jesus-mata/academy-go-q12021/infrastructure/config"
	"github.com/jesus-mata/academy-go-q12021/infrastructure/handler"
	"github.com/jesus-mata/academy-go-q12021/infrastructure/routes"
	"github.com/jesus-mata/academy-go-q12021/interfaces/repository"
	"github.com/jesus-mata/academy-go-q12021/usecase/interactors"
)

func main() {
	logger := log.New(os.Stdout, "News API ", log.LstdFlags|log.Lshortfile)

	router := mux.NewRouter().StrictSlash(true)

	config.ReadConfig()
	port := config.GetString("app.port")
	appName := config.GetString("app.name")
	addr := fmt.Sprint(":", port)

	logger.Println("Starting server on", addr)
	srv := infrastructure.NewServer(router, addr)

	nh := provideNewsHandler(logger)
	routes.SetupRoutes(router, nh)

	logger.Println(appName, "started and listening requests at", addr)

	err := srv.ListenAndServe()
	if err != nil {
		logger.Fatalf("Server failed to start: %v", err)
	}

}

func provideNewsHandler(logger *log.Logger) *handler.NewsHandlers {

	csvReader := infrastructure.NewCsvReader("./resources/data.csv", logger)
	newsRepository := repository.NewNewsArticleRepository(csvReader, logger)

	newsInteractor := interactors.NewNewsArticlesInteractor(newsRepository)

	nh := handler.NewNewsHandlers(newsInteractor)
	return nh
}
