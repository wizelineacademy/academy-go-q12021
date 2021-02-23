package models

// Pokemon Type
type Pokemon struct {
	ID   int    `json:ID`
	Name string `json:Name`
	URL  string `json:Url`
}

type SinglePokeExternal struct {
	Name string `json:name`
	URL  string `json:url`
}

type PokemonExternal struct {
	Count    int                  `json:count`
	Next     string               `json:next`
	Previous string               `json:previous`
	Results  []SinglePokeExternal `json:results`
}
