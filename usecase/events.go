package usecase

import (
	"fmt"

	"github.com/javiertlopez/golang-bootcamp-2020/model"
	"github.com/javiertlopez/golang-bootcamp-2020/repository"
)

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

type events struct {
	eventRepo       repository.EventRepository
	reservationRepo repository.ReservationRepository
}

// NewEventUseCase returns the usecase implementation
func NewEventUseCase(
	eventRepo repository.EventRepository,
	reservationRepo repository.ReservationRepository,
) Events {
	return &events{
		eventRepo,
		reservationRepo,
	}
}

func (e *events) Create(event model.Event) (model.Event, error) {
	return model.Event{}, fmt.Errorf("not implemented")
}

func (e *events) GetByID(id string) (model.Event, error) {
	return model.Event{}, fmt.Errorf("not implemented")
}

func (e *events) GetAll() ([]model.Event, error) {
	return nil, fmt.Errorf("not implemented")
}

func (e *events) Update(event model.Event) (model.Event, error) {
	return model.Event{}, fmt.Errorf("not implemented")
}

func (e *events) Delete(id string) error {
	return fmt.Errorf("not implemented")
}

func (e *events) AddReservations(id string, reservations []model.Reservation) ([]model.Reservation, error) {
	return nil, fmt.Errorf("not implemented")
}

func (e *events) GetReservations(id string) ([]model.Reservation, error) {
	return nil, fmt.Errorf("not implemented")
}
