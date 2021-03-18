package registry

import (
	"bootcamp/interface/controller"
	"database/sql"
)

type registry struct {
	db *sql.DB
}

type Registry interface {
	NewAppController() controller.AppController
}

func NewRegistry(db *sql.DB) Registry {
	return &registry{db}
}

func (r *registry) NewAppController() controller.AppController {
	return controller.AppController {
		Item: r.NewItemController(),
		Joke: r.NewJokeController(),
	}
}