package handler

import (
	"encoding/json"
	"net/http"

	"github.com/javiertlopez/golang-bootcamp-2020/model"

	"github.com/gorilla/mux"
)

// addEventHandler adds the handler to the mux router
func (h *Handler) addEventHandler(r *mux.Router) {
	r.HandleFunc("/events", h.CreateEvent).Methods("POST")
	r.HandleFunc("/events/{id}", h.GetEventByID).Methods("GET")
	r.HandleFunc("/events/{id}/reservations", h.GetReservations).Methods("GET")
}

// CreateEvent handler
func (h *Handler) CreateEvent(w http.ResponseWriter, r *http.Request) {
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

	// Step 1. Validate reservations
	if event.Reservations != nil {
		err := h.validateReservation(event.Reservations)
		if err != nil {
			JSONResponse(
				w, http.StatusUnprocessableEntity,
				Response{
					Message: err.Error(),
					Status:  http.StatusUnprocessableEntity,
				},
			)
			return
		}
	}

	// Step 2. Create event
	response, err := h.events.Create(event)
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

	// Step 3. Get TotalFee
	response.TotalFee = response.CalculateTotalFee()

	JSONResponse(
		w,
		http.StatusCreated,
		response,
	)
}

// GetEventByID handler
func (h *Handler) GetEventByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	response, err := h.events.GetByID(id)
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
	reservations, err := h.events.GetReservations(id)
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

	// Step 2. Get TotalFee
	response.Reservations = reservations
	response.TotalFee = response.CalculateTotalFee()

	JSONResponse(
		w,
		http.StatusCreated,
		response,
	)
}
