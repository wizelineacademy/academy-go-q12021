package api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
)

// Health is the healtcheck endpoint
func (s *PokeAPI) Health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "OK")
}

// GetPokeByID parses the id from the query param and returns the
// selected pokemon from the service
func (s *PokeAPI) GetPokeByID(w http.ResponseWriter, r *http.Request) {
	pokeIDraw := r.URL.Query().Get("id")
	pokeID, err := strconv.Atoi(pokeIDraw)
	if err != nil {
		log.Printf("wrong API argument: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, "'id' parameter required")
		return
	}

	w.Header().Set("Content-Type", "application/json")

	p, err := s.PokeService.GetPokeByID(pokeID)
	if err != nil {
		log.Printf("error fetching pokemon: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, "error fetching pokemon")
		return
	}
	pl := &PokeWrapper{
		ID:          pokeID,
		PokemonData: p,
	}
	log.Printf("serving pokemon id: %d", pokeID)
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(pl); err != nil {
		log.Fatalln("cant encode")
	}
}
