package models

type CardBack struct {
	ID           int    `json:"id"`
	SortCategory int    `json:"sortCategory"`
	Text         string `json:"text"`
	Name         string `json:"name"`
	Image        string `json:"image"`
	Slug         string `json:"slug"`
}
