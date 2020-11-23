package models

import (
	"encoding/json"
	"io"
	"time"
)

// Champion defines a champion attributes
type Champion struct {
	ID          int       `json:"-"`
	Name        string    `json:"name"`
	Lore        string    `json:"lore"`
	DateCreated time.Time `json:"created_at"`
}

// ToJSON encodes a Champion struct to JSON
func (c *Champion) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(c)
}

// FromJSON decodes a JSON object to a Champion struct
func (c *Champion) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(c)
}
