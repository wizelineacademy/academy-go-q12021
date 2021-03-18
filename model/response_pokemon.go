package model

type ResponsePokemon struct {
    Error   string      `json:"error,omitempty"`
    Result  []Pokemon   `json:"result"`
    Total   int         `json:"total"`
}
