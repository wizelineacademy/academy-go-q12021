package errors

import (
    "fmt"
)

type PokemonNotFoundError struct {
    Id      int
    ErrMsg  error
}

func (e PokemonNotFoundError) Error() string {
    return fmt.Sprintf("Pokemon with index: %v not found", e.Id)
}

func (e PokemonNotFoundError) Unwrap() error {
    return e.ErrMsg
}
