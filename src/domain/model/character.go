package model

//Character - character model
type Character struct {
	ID       int      `json:"id"`
	Name     string   `json:"name"`
	Status   string   `json:"status"`
	Species  string   `json:"species"`
	Type     string   `json:"type"`
	Gender   string   `json:"gender"`
	Origin   Nested   `json:"origin"`
	Location Nested   `json:"location"`
	Image    string   `json:"image"`
	Episodes []string `json:"episode"`
}

//Nested - nested data of interest
type Nested struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
