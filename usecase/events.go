package usecase

import "github.com/javiertlopez/golang-bootcamp-2020/model"

// Events interface
type Events interface {
	// Event related
	Create(event model.Event) (model.Event, error)
	GetByID(id string) (model.Event, error)
	GetAll() ([]model.Event, error)
	Update(event model.Event) (model.Event, error)
	Delete(id string) error

	// Reservation related. For CSV use.
	AddReservations(id string, reservations []model.Reservation) ([]model.Reservation, error)
	GetReservations(id string) ([]model.Reservation, error)

	// Do I need to pass the context?
}
