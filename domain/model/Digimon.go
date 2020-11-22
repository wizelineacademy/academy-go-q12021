package model

// Digimon is the structure to contain digimon data.
type Digimon struct {
	Name  string `gorm:"primary_key" json:"name"`
	Level string `json:"level"`
	Image string `json:"img"`
}
