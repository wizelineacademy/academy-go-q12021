package pokemon

import (
	"database/sql"
	"fmt"
	"strings"
)

// Pokemon model for the pokemon
type Pokemon struct {
	Number   int
	Name     string
	Type1    sql.NullString
	Type2    sql.NullString
	HeightM  float64
	WeightKg float64
	Male     float64
	Female   float64
	CaptRate float64
	HP       int
	Attack   int
	Defense  int
	Special  int
	Speed    int
}

// Show returns the pokemon name and type
func (p *Pokemon) Show() string {
	var types []string

	for _, val := range []sql.NullString{p.Type1, p.Type2} {
		if val.Valid {
			types = append(types, strings.Title(val.String))
		}
	}

	return fmt.Sprintf("{%03d - %v: %v}", p.Number, p.Name, strings.Join(types, " "))
}

// Describe shows detailed information on the pokemon
func (p *Pokemon) Describe() string {

	return fmt.Sprintf(
		"%v\n  HP:  %4d\n  Atk: %4d\n  Def: %4d\n  Spe: %4d\n  Spd: %4d",
		p.Show(),
		p.HP,
		p.Attack,
		p.Defense,
		p.Special,
		p.Speed)
}
