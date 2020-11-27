package usecase

import (
	"time"

	"github.com/javiertlopez/golang-bootcamp-2020/model"
	"github.com/javiertlopez/golang-bootcamp-2020/repository"

	guuid "github.com/google/uuid"
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

// Create a new event. Reservation is optional.
func (e *events) Create(event model.Event) (model.Event, error) {
	// Step 0. Let's create a UUID
	uuid := guuid.New().String()
	event.ID = uuid

	// Step 0.1. Now!
	now := time.Now()
	event.CreatedAt = &now
	event.UpdatedAt = &now

	// Step 1. Try to store an event
	event, err := e.eventRepo.Create(event)
	if err != nil {
		return model.Event{}, err
	}

	// Step 2. If there are reservations, store them
	if event.Reservations != nil {
		event.Reservations, err = e.AddReservations(uuid, event.Reservations)

		if err != nil {
			return model.Event{}, err
		}
	}

	return event, nil
}

// GetByID returns an event and its reservations
func (e *events) GetByID(id string) (model.Event, error) {
	// Step 1. Get event
	event, err := e.eventRepo.GetByID(id)
	if err != nil {
		return model.Event{}, err
	}

	// Step 2. Get reservations
	reservations, err := e.GetReservations(id)
	if err != nil {
		return model.Event{}, err
	}

	event.Reservations = reservations

	return event, nil
}

// GetAll returns all events. Doesn't return reservations, too expensive
func (e *events) GetAll() ([]model.Event, error) {
	events, err := e.eventRepo.GetAll()
	if err != nil {
		return nil, err
	}

	return events, nil
}

// Update an event. Where should I prevent certain fields from being updated?
func (e *events) Update(event model.Event) (model.Event, error) {
	event, err := e.eventRepo.Update(event)
	if err != nil {
		return model.Event{}, err
	}

	return event, nil
}

// Delete an event. This should be an idempotent request
// what should I return if the event has already been deleted?
func (e *events) Delete(id string) error {
	err := e.eventRepo.Delete(id)
	if err != nil {
		return nil
	}

	return err
}

// AddReservations stores reservations, and adds an ID per reservation.
func (e *events) AddReservations(id string, reservations []model.Reservation) ([]model.Reservation, error) {
	for _, res := range reservations {
		// Step 0. Let's create a UUID
		uuid := guuid.New().String()
		res.ID = uuid

		// Step 1. Store the reservation
		_, err := e.reservationRepo.Create(id, res)
		if err != nil {
			return nil, err
		}
	}

	return reservations, nil
}

// GetReservations returns all reservations for a given event
func (e *events) GetReservations(id string) ([]model.Reservation, error) {
	reservations, err := e.reservationRepo.GetByEventID(id)
	if err != nil {
		return nil, err
	}

	return reservations, nil
}
