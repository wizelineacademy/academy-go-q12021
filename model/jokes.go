package model

type Joke struct {
	ID        int    `json:id`
	Setup     string `json:setup`
	Punchline string `json:punchline`
}
