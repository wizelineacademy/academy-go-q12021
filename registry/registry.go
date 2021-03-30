package registry

import (
	"bootcamp/interface/controller"
	"database/sql"
)

// registry struct for SQL DB
type registry struct {
	db *sql.DB
}

// Registry interface for app controllers
type Registry interface {
	NewAppController() controller.AppController
}

// NewRegistry func to setup SQL interface
func NewRegistry(db *sql.DB) Registry {
	return &registry{db}
}

// NewAppController func for Items and Jokes
func (r *registry) NewAppController() controller.AppController {
	return controller.AppController {
		Item: r.NewItemController(),
		Joke: r.NewJokeController(),
	}
}