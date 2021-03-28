package model

// Pokemon Type
type Pokemon struct {
	ID   int    `json:ID,omitempty`
	Name string `json:Name,omitempty`
	URL  string `json:Url,omitempty`
}

type SinglePokeExternal struct {
	Name string `json:name,omitempty`
	URL  string `json:url,omitempty`
}

type PokemonExternal struct {
	Count    int                  `json:count`
	Next     string               `json:next`
	Previous string               `json:previous`
	Results  []SinglePokeExternal `json:results`
}

type TypeNumberFilter struct {
	Even string
	Odd  string
}
