package config

import (
	"net/http"

	"github.com/wizelineacademy/golang-bootcamp-2020/domain/repositories"

	"github.com/wizelineacademy/golang-bootcamp-2020/interface/handlers"

	"github.com/gorilla/mux"
)

// InitRoutes initialize the mux routes
func (c *Config) InitRoutes() *mux.Router {
	// Champions handler
	ch := handlers.NewChampionHandler(c.InfoLog, c.ErrorLog, repositories.NewChampRepo(c.DB))

	// Gorilla serve mux
	sm := mux.NewRouter()

	// Gorilla routes
	// Get
	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/champions", ch.GetChamps)
	getRouter.HandleFunc("/champions/{id:[0-9]+$}", ch.GetChamp)

	// Post
	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/champion", ch.AddChamp)

	//postRouter.Use(ch.MiddlewareUserValidation)

	// Put
	// putRouter := sm.Methods(http.MethodPut).Subrouter()
	// putRouter.HandleFunc("/{id:[0-9]+}", ch.UpdateChamp)

	return sm
}
