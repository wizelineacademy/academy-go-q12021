package controller

import (
	"net/http"

	"github.com/javiertlopez/golang-bootcamp-2020/usecase"
)

// EventController handles the HTTP requests
type EventController interface {
	CreateEvent(w http.ResponseWriter, r *http.Request)
	GetEventByID(w http.ResponseWriter, r *http.Request)
	GetReservations(w http.ResponseWriter, r *http.Request)
}

// eventController struct holds the usecase
type eventController struct {
	events usecase.Events
}

// NewEventController returns an EventController
func NewEventController(events usecase.Events) EventController {
	return &eventController{
		events,
	}
}
