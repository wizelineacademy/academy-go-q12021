package repository

import "github.com/javiertlopez/golang-bootcamp-2020/model"

// EventRepository interface
type EventRepository interface {
	// Event related
	Create(event model.Event) (model.Event, error)
	GetByID(id string) (model.Event, error)
	GetAll() ([]model.Event, error)
	Update(event model.Event) (model.Event, error)
	Delete(id string) error
}

// ReservationRepository interface
type ReservationRepository interface {
	// Reservation related
	Create(eventID string, reservation model.Reservation) (model.Reservation, error)

	// Related to Event
	GetByEventID(id string) ([]model.Reservation, error)
}
