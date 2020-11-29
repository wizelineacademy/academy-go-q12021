package router

import "github.com/gorilla/mux"

// addEventController setup the router with the event controller
func (r *Router) addEventController(router *mux.Router) {
	router.HandleFunc("/events", r.events.CreateEvent).Methods("POST")
	router.HandleFunc("/events/{id}", r.events.GetEventByID).Methods("GET")
	router.HandleFunc("/events/{id}/reservations", r.events.GetReservations).Methods("GET")
}
