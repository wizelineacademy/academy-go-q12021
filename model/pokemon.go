package model

import (
	"fmt"
)

type Pokemon struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func (p Pokemon) String() string {
	return fmt.Sprintf("\n{\n\tID: %v,\n\tName: %v\n}\n", p.Id, p.Name)
}
