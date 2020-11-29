package handler

import (
	"fmt"
	"net/http"

	"github.com/javiertlopez/golang-bootcamp-2020/model"

	"github.com/gorilla/mux"
)

// GetReservations handler
func (h *Handler) GetReservations(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	_, err := h.events.GetByID(id)
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

	JSONResponse(
		w,
		http.StatusOK,
		reservations,
	)
}

func (h *Handler) validateReservation(reservations []model.Reservation) error {
	// Step 1. Validate fees
	if reservations != nil {
		for _, res := range reservations {
			// 1.1 Adults should be at least 1
			if res.Adults < 1 {
				return fmt.Errorf("res.Adults < 1")
			}

			// 1.2 AdultFee should be more than 0
			if res.AdultFee <= 0 {
				return fmt.Errorf("res.AdultFee <= 0")
			}

			// 1.3 Minors should be more or equal to 0
			if res.Minors < 0 {
				return fmt.Errorf("res.Minors <= 0")
			}

			// 1.4 if Minors, MinorFee should be more than 0
			if res.Minors > 0 && res.MinorFee <= 0 {
				return fmt.Errorf("res.Minors > 0 && res.MinorFee <= 0")
			}
		}
	}

	return nil
}
