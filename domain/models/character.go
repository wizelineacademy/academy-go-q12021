package models

type Character struct {
	ID       int       `json:"id,omitempty" csv:"id"`
	Name     string    `json:"name,omitempty" csv:"name"`
	Status   string    `json:"status,omitempty" csv:"status"`
	Species  string    `json:"species,omitempty" csv:"species"`
	Type     string    `json:"type,omitempty" csv:"-"`
	Location *Location `json:"location,omitempty" csv:"-"`
	Origin   *Origin   `json:"origin,omitempty" csv:"-"`
	Error    string    `json:"error,omitempty" csv:"-"`
}

type Location struct {
	Name string `json:"name,omitempty"`
	URL  string `json:"url,omitempty"`
}

type Origin struct {
	Name string `json:"name,omitempty"`
	URL  string `json:"url,omitempty"`
}
