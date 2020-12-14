package controller

import (
	"encoding/json"
	"net/http"

	"github.com/javiertlopez/golang-bootcamp-2020/errorcodes"
	"github.com/javiertlopez/golang-bootcamp-2020/model"

	"github.com/gorilla/mux"
)

// CreateEvent handler
func (e *eventController) CreateEvent(w http.ResponseWriter, r *http.Request) {
	var event model.Event
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&event); err != nil {
		JSONResponse(
			w, http.StatusBadRequest,
			Response{
				Message: "Bad request",
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
				Message: "Internal server error",
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

// GetEventByID handler
func (e *eventController) GetEventByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	response, err := e.events.GetByID(id)
	if err != nil {
		if err == errorcodes.ErrEventNotFound {
			JSONResponse(
				w, http.StatusNotFound,
				Response{
					Message: "Not found",
					Status:  http.StatusNotFound,
				},
			)
			return
		}

		JSONResponse(
			w, http.StatusInternalServerError,
			Response{
				Message: "Internal server error",
				Status:  http.StatusInternalServerError,
			},
		)
		return
	}

	JSONResponse(
		w,
		http.StatusOK,
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
				Message: "Not found",
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
				Message: "Internal server error",
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
