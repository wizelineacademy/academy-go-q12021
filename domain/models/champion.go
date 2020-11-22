package models

import (
	"encoding/json"
	"errors"
	"io"
	"time"
)

var ErrNoRecord = errors.New("models: no matching record found")

// Champion defines a champion attributes
type Champion struct {
	ID          int       `json:"-"`
	Name        string    `json:"name"`
	Lore        string    `json:"lore"`
	DateCreated time.Time `json:"created_at"`
}

// ChampionRepository defines the interface to use with a Champion
type ChampionRepository interface {
	GetSingle(id int) (*Champion, error)
	GetMultiple() ([]*Champion, error)
}

func (c *Champion) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(c)
}
