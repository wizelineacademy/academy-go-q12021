package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/javiertlopez/golang-bootcamp-2020/model"
)

// CreateEvent handler
func (e *eventController) CreateEvent(w http.ResponseWriter, r *http.Request) {
	var event model.Event
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&event); err != nil {
		JSONResponse(
			w, http.StatusBadRequest,
			Response{
				Message: err.Error(),
				Status:  http.StatusBadRequest,
			},
		)
		return
	}
	defer r.Body.Close()

	// Step 2. Create event
	response, err := e.events.Create(event)
	if err != nil {
		JSONResponse(
			w, http.StatusInternalServerError,
			Response{
				Message: err.Error(),
				Status:  http.StatusBadRequest,
			},
		)

		return
	}

	JSONResponse(
		w,
		http.StatusCreated,
		response,
	)
}

// GetEventByID handler
func (e *eventController) GetEventByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	response, err := e.events.GetByID(id)
	if err != nil {
		JSONResponse(
			w, http.StatusNotFound,
			Response{
				Message: err.Error(),
				Status:  http.StatusNotFound,
			},
		)
		return
	}

	// Step 1. Get reservations
	response.Reservations, err = e.events.GetReservations(id)
	if err != nil {
		JSONResponse(
			w, http.StatusInternalServerError,
			Response{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
			},
		)
		return
	}

	JSONResponse(
		w,
		http.StatusCreated,
		response,
	)
}

// GetReservations handler
func (e *eventController) GetReservations(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	_, err := e.events.GetByID(id)
	if err != nil {
		JSONResponse(
			w, http.StatusNotFound,
			Response{
				Message: err.Error(),
				Status:  http.StatusNotFound,
			},
		)
		return
	}

	// Step 1. Get reservations
	reservations, err := e.events.GetReservations(id)
	if err != nil {
		JSONResponse(
			w, http.StatusInternalServerError,
			Response{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
			},
		)
		return
	}

	JSONResponse(
		w,
		http.StatusOK,
		reservations,
	)
}
