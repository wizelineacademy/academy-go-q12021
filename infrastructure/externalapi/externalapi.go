package externalapi

import (
  "fmt"
  "net/http"
  "io/ioutil"
  "encoding/json"
  "github.com/halarcon-wizeline/academy-go-q12021/domain"
)

// Retrieve pokemons from external url
func GetPokemons () ([]domain.Pokemon, error) {

  var pokemons []domain.Pokemon

  response, err := http.Get("https://pokeapi.co/api/v2/pokemon")
  if err != nil {
    fmt.Printf("The HTTP request failed with error %s", err)
    return pokemons, err
  }

  data, _ := ioutil.ReadAll(response.Body)
  // fmt.Println(string(data))
  defer response.Body.Close()

  pokemons = parsePokemonsWithStructs(data)

  return pokemons, nil
}

// Read elements with structs
func parsePokemonsWithStructs (data []byte) ([]domain.Pokemon) {

  var pokemons []domain.Pokemon
  var pokemonRequest domain.PokemonApiRequest

  json.Unmarshal([]byte(data), &pokemonRequest)
  // fmt.Println(pokemonRequest.Results[1].ID)
  // fmt.Println(pokemonRequest.Results[1].Name)

  for key, element := range pokemonRequest.Results {
    // pokemon := domain.NewPokemon(key, element.Name)
    pokemon := domain.Pokemon {ID:key, Name:element.Name}
    pokemons = append(pokemons, pokemon)
  }
  return pokemons
}

/*
func parsePokemonsWithoutStructs (data []byte) ([]domain.Pokemon) {

  var pokemons []domain.Pokemon

  // Read elements without structs
  var dataJson map[string]interface{}
  json.Unmarshal([]byte(data), &dataJson)
  // fmt.Println(dataJson["results"])

  pokemons := dataJson["results"].([]interface{})

  i := 0
  for _, value := range pokemons {
   element := value.(map[string]interface{})
   fmt.Println(element["name"])
   pokemon := domain.NewPokemon(i, string(element["name"]))
   i := i + 1
   // fmt.Println(pokemon.Name)
  }
}
*/
