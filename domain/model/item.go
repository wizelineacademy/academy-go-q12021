package model

// Item represents the CSV file structure in a model struct
type Item struct {
	ID uint `gorm:"primary_key" json:"id"`
	Name string `json:"name"`
}

// TableName function returns the CSV filename,
// the CSVQ driver adds the extension file ".csv".
func (Item) TableName() string { return "items" }
