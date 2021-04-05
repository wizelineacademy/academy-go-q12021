package models

//Pokemon_api is a struct for the response of pokemon api
type Pokemon_api struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	URL  string `json:"url"`
}
