package models

//Pokemon is a struct to get the information from csv file
type Pokemon struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
