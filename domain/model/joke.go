package model

// Joke represents the CSV file structure and API response in a model struct
type Joke struct {
	ID     string `gorm:"primary_key" json:"id"`
	Joke   string `json:"joke"`
	Status int    `json:"status"`
}

// TableName function returns the CSV filename,
// the CSVQ driver adds the extension file ".csv".
func (Joke) TableName() string { return "jokes" }
