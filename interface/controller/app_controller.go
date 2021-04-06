package controller

// AppController contains all controllers in the app
type AppController struct {
	Item interface { ItemController }
	Joke interface { JokeController }
}
