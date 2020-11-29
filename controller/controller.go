package controller

import (
	"net/http"
)

// EventController handles the HTTP requests
type EventController interface {
	CreateEvent(w http.ResponseWriter, r *http.Request)
	GetEventByID(w http.ResponseWriter, r *http.Request)
	GetReservations(w http.ResponseWriter, r *http.Request)
}
