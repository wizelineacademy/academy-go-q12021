package model

import (
	"fmt"
)

type Pokemon struct {
	Id int `json:"id"`
	Name string `json:"name"`
}

func (p Pokemon) String() string {
	return fmt.Sprintf("ID\t\t\t|Name\n%v\t\t\t|%v", p.Id, p.Name)
}
