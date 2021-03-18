package controller

import (
    "fmt"
    "log"
    "net/http"
    "strconv"
    "encoding/json"


    "pokedex/service"
    "pokedex/model"

    "github.com/gorilla/mux"
)

/* Used to initialize Pokedex Services */
var pokedex = service.Pokedex(make(map[int]model.Pokemon))

/* Obtains pokemon by id and returns it in JSON string format */
func get_pokemon_by_id(w http.ResponseWriter, r *http.Request) {
    query_id, ok := r.URL.Query()["id"]
    response := model.ResponsePokemon{Error: "No available data found"}

    if ok {
        id, err := strconv.Atoi(query_id[0])
        if err == nil {
            result := make([]model.Pokemon, 1)

            /* Check if pokemon was found */
            pokemon, err := pokedex.GetPokemonById(id)
            if err == nil {
                result[0] = pokemon
                response = model.ResponsePokemon{Result: result, Total: len(result)}

            }
        }
    }

    json.NewEncoder(w).Encode(response)
    return
}

/* Homepage */
func home(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Wizeline - Go Bootcamp 2021 - @hugoaguirre")
}

func InitPokedexApp() {
    /* Init Pokedex */
    pokedex.Init()

    /* Routers init */
    pokedex_router := mux.NewRouter().StrictSlash(true)

    pokedex_router.HandleFunc("/", home)
    pokedex_router.HandleFunc("/pokemon", get_pokemon_by_id)

    log.Fatal(http.ListenAndServe("localhost:8080", pokedex_router))
}
