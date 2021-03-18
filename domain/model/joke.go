package model

type Joke struct {
	ID     string "json:id"
	Joke   string "json:joke"
	Status int    "json:status"
}
