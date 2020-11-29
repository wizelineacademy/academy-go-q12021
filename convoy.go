package main

import (
	"context"
	"time"

	"github.com/javiertlopez/golang-bootcamp-2020/controller"
	"github.com/javiertlopez/golang-bootcamp-2020/repository/axiom"
	"github.com/javiertlopez/golang-bootcamp-2020/router"
	"github.com/javiertlopez/golang-bootcamp-2020/usecase"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	mongoTimeout = 15 * time.Second
)

// App holds the handler, and logger
type App struct {
	logger *logrus.Logger
	router *mux.Router
	config AppConfig
}

// AppConfig struct with configuration variables
type AppConfig struct {
	MongoDB  string
	MongoURI string
}

// New returns an App
func New(config AppConfig, logger *logrus.Logger) App {
	// Set client options
	clientOptions := options.Client().ApplyURI(config.MongoURI)

	// Context with timeout for establish connection with Mongo Atlas
	ctx, cancel := context.WithTimeout(context.Background(), mongoTimeout)
	defer cancel()

	// Connect to Mongo Atlas
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		logger.Fatal(err)
	}

	// Init eventsRepository
	eventsRepo := axiom.NewEventsRepo(config.MongoDB, client)

	// Init reservationsRepository
	reservationsRepo := axiom.NewReservationRepo(config.MongoDB, client)

	// Init usecase
	events := usecase.NewEventUseCase(eventsRepo, reservationsRepo)

	// Init controller
	controller := controller.NewEventController(events)

	// Setup router
	router := router.New(controller)

	return App{
		logger,
		router.Router(),
		config,
	}
}

// Router returns *mux.Router
func (a *App) Router() *mux.Router {
	return a.router
}
