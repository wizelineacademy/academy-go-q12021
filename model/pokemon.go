package model

import (
    "fmt"
)

type Pokemon struct {
    Id int      `json:"id"`
    Name string `json:"name"`
}

func (pokemon Pokemon) String() string {
    return fmt.Sprintf("ID: %v, Name: %v", pokemon.Id, pokemon.Name )
}
