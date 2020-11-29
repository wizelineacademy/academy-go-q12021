package router

import (
	"github.com/javiertlopez/golang-bootcamp-2020/controller"

	"github.com/gorilla/mux"
)

// Router struct holds the event controller
type Router struct {
	events controller.EventController
}

// New returns a Router
func New(events controller.EventController) Router {
	return Router{
		events,
	}
}

// Router returns a *mux.Router
func (r *Router) Router() *mux.Router {
	// Create a Gorilla Mux router
	router := mux.NewRouter()

	// addEventController
	r.addEventController(router)

	return router
}
