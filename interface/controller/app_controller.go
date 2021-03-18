package controller

type AppController struct {
	Item interface { ItemController }
	Joke interface { JokeController }
}
