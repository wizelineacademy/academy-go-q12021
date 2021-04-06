package handlers

import "net/http"

//GetPokemonsConcurrency is the handler to get the pokemons with concurrency
func (p *Pokemons) GetPokemonsConcurrency(w http.ResponseWriter, r *http.Request) {
	odd := r.FormValue("odd")
	quantity := r.FormValue("quantity")
	nWorkers := r.FormValue("numberWorkers")

	if odd == "" || quantity == "" || nWorkers == "" {
		p.logger.Error("Missing query params")
		p.render.JSON(w, http.StatusBadRequest, "Missing query params")
		return
	}

	tmp, err := p.useCase.GetPokemonsConcurrency(odd, quantity, nWorkers)
	if err != nil {
		p.logger.Error(err)
		p.render.JSON(w, http.StatusInternalServerError, tmp)
	}
	p.render.JSON(w, http.StatusOK, tmp)
}
