package model

//Episode - episode model
type Episode struct {
	ID         int      `json:"id"`
	Name       string   `json:"name"`
	AirDate    string   `json:"air_date"`
	Characters []string `json:"characters"`
}
