package service

import (
    "os"
    "fmt"
    "strings"
    "strconv"

    "pokedex/config"
    "pokedex/model"
    "pokedex/model/errors"
    "pokedex/data"
)

/* This map will act as an in-RAM Pokedex */
type Pokedex map[int]model.Pokemon

/* Parses data read from CSV into a Pokedex map so it can be consumed later */
func (pokedex Pokedex) Init() {
    path, err := os.Getwd()
    if err != nil {
        fmt.Println(err)
    }

    data_source := data.CsvSource(path + "/" + config.PokedexConfig.CSV_SOURCE)
    items, err := data_source.InitCsvSource()

    if err != nil {
        return
    }

    /* Data obtained from CsvSource will come in string format
    * e.g.:  4 charmander
    */
    for _, line := range items {
        id := strings.Trim(line[0], " ")
        name := strings.Trim(line[1], " ")

        if len(name) == 0 {
            fmt.Printf("invalid name length found for pokemon index: %v (empty name)", id)
            continue
        }

        id_conv, err := strconv.Atoi(id)
        if err != nil {
            fmt.Printf("invalid index found for pokemon: %v (not an integer)", name)
            continue
        }

        pokedex[id_conv] = model.Pokemon{Id: id_conv, Name: name}
    }

    if len(pokedex) == 0 {
        pokedex = nil
    }
}

/* Obtains a pokemon based on its id */
func (pokedex Pokedex) GetPokemonById(id int) (model.Pokemon, error) {
    if pokedex == nil {
        no_data_err := errors.NoDataError{}
        return model.Pokemon{}, no_data_err
    }

    pokemon, ok := pokedex[id]
    if ok {
        return pokemon, nil
    }

    not_found_err := errors.PokemonNotFoundError{Id: id}
    return model.Pokemon{}, not_found_err
}
