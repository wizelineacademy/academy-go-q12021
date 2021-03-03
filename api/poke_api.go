package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/wizelineacademy/academy-go-q12021/poke"
)

type PokeApi struct {
	Router      *http.ServeMux
	context     context.Context
	port        int
	server      *http.Server
	PokeService *poke.PokeService
}

type PokeWrapper struct {
	ID          int            `json:"id"`
	PokemonData []poke.Pokemon `json:"pokemon_data"`
}

func NewPokeApi(ctx context.Context, port int, pokeService *poke.PokeService) *PokeApi {

	service := PokeApi{
		Router:      http.NewServeMux(),
		context:     ctx,
		port:        port,
		PokeService: pokeService,
	}
	service.setupRouter()
	return &service
}

func (s *PokeApi) setupRouter() {
	s.Router.HandleFunc("/", s.Health)
	s.Router.HandleFunc("/getPoke/", s.GetPokeById)
}

func (s *PokeApi) StartServer() error {
	p := fmt.Sprintf(":%d", s.port)
	s.server = &http.Server{
		Addr:    p,
		Handler: s.Router,
	}
	return s.server.ListenAndServe()
}
