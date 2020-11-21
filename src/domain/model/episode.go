package model

type Episode struct {
	Id         int      `json:"id"`
	Name       string   `json:"name"`
	AirDate    string   `json:"air_date"`
	Characters []string `json:"characters"`
}
