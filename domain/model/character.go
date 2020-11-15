package model

type Character struct {
	Id       int      `json:"id"`
	Name     string   `json:"name"`
	Status   string   `json:"status"`
	Species  string   `json:"species"`
	Type     string   `json:"type"`
	Gender   string   `json:"gender"`
	Origin   Location `json:"origin"`
	Location Location `json:"location"`
	Image    string   `json:"image"`
	Episodes []string `json:"episode"`
}
