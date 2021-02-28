package controller

// AppController struct
type AppController struct {
	Pokemon interface{ PokemonController }
}
