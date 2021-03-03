package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/wizelineacademy/academy-go-q12021/poke"
)

// PokeAPI is the main API object
type PokeAPI struct {
	Router      *http.ServeMux
	context     context.Context
	port        int
	server      *http.Server
	PokeService *poke.Service
}

// PokeWrapper holds the response JSON struct
type PokeWrapper struct {
	ID          int            `json:"id"`
	PokemonData []poke.Pokemon `json:"pokemon_data"`
}

// NewPokeAPI creates and setups an instance of the api server
func NewPokeAPI(ctx context.Context, port int, pokeService *poke.Service) *PokeAPI {

	service := PokeAPI{
		Router:      http.NewServeMux(),
		context:     ctx,
		port:        port,
		PokeService: pokeService,
	}
	service.setupRouter()
	return &service
}

func (s *PokeAPI) setupRouter() {

	s.Router.HandleFunc("/", s.Health)
	s.Router.HandleFunc("/getPoke/", s.GetPokeByID)
}

// StartServer initialize the server at the specified port
func (s *PokeAPI) StartServer() error {
	p := fmt.Sprintf(":%d", s.port)
	s.server = &http.Server{
		Addr:    p,
		Handler: s.Router,
	}
	return s.server.ListenAndServe()
}
