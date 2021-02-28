package model

// Pokemon data struct
type Pokemon struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	URL  string `json:"url"`
}

// TableName returns the name of the table
func (Pokemon) TableName() string { return "pokemons" }
