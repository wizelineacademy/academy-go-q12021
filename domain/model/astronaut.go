package model

// Astronaut is the main structure for collecting data from the CSV
type Astronaut struct {
	Name     string `json:"name"`
	FlightHr string `json:"flightHr"`
}
