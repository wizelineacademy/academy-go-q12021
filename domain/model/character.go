package model

type Character struct {
	Id       int      `json:"id"`
	Name     string   `json:"name"`
	Status   string   `json:"status"`
	Species  string   `json:"species"`
	Type     string   `json:"type"`
	Gender   string   `json:"gender"`
	Origin   nested   `json:"origin"`
	Location nested   `json:"location"`
	Image    string   `json:"image"`
	Episodes []string `json:"episode"`
}

type nested struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}
