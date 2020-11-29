package model

import "time"

// Event holds information related to an event
type Event struct {
	ID            string     `json:"id"`
	Description   string     `json:"description"`
	Type          string     `json:"type"` // how can I predefine values here?
	Status        string     `json:"status"`
	CreatedAt     *time.Time `json:"created_at"`
	UpdatedAt     *time.Time `json:"updated_at"`
	EventDate     *time.Time `json:"event_date"`     // should I drop 'event_'?
	EventLocation string     `json:"event_location"` // should I drop 'event_'?

	// Customer information
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Email string `json:"email"`

	// Guests
	Reservations []Reservation `json:"reservations,omitempty"`
	TotalFee     float64       `json:"total_fee,omitempty"`
}

// CalculateTotalFee returns the total cost of the event guests
func (e *Event) CalculateTotalFee() {
	var total float64

	for _, reservation := range e.Reservations {
		total += reservation.totalFee()
	}

	e.TotalFee = total
}

// filterReservations returns reservations filtered by status
func (e *Event) filterReservations(status string) []Reservation {
	var list []Reservation

	for _, reservation := range e.Reservations {
		if e.Status == status {
			list = append(list, reservation)
		}
	}

	return list
}
