package model

import "time"

// Reservation holds information for a single hotel reservation
type Reservation struct {
	ID        string     `json:"id"`
	Status    string     `json:"status"` // how can I predefine values here?
	Plan      string     `json:"plan"`   // how can I predefine values here?
	Adults    int        `json:"adults"`
	Minors    int        `json:"minors"`
	AdultFee  float64    `json:"adult_fee"`
	MinorFee  float64    `json:"minor_fee"`
	Arrival   *time.Time `json:"arrival"`
	Departure *time.Time `json:"departure"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`

	// Customer information
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Email string `json:"email"`
}

// totalFee returns the total cost of the reservation
func (r *Reservation) totalFee() float64 {
	// where should I validate that the fees are above 0?
	// we have to calculate the total fee per night
	cost := (float64(r.Adults) * r.AdultFee) + (float64(r.Minors) * r.MinorFee)

	// we have to calculate the total fee
	total := float64(r.nights()) * cost

	return total
}

// nights returns the total days of stay
func (r *Reservation) nights() int {
	// filter out hours, minutes and seconds
	arrival := time.Date(r.Arrival.Year(), r.Arrival.Month(), r.Arrival.Day(), 0, 0, 0, 0, time.UTC)
	departure := time.Date(r.Departure.Year(), r.Departure.Month(), r.Departure.Day(), 0, 0, 0, 0, time.UTC)

	return int(departure.Sub(arrival).Hours() / 24)
}
