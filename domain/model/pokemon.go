package model

import "strconv"

type Pokemon struct {
	ID         int
	Name       string
	Type       string
	Subtype    string
	Total      int
	HP         int
	Attack     int
	Defense    int
	SpAtk      int
	SpDef      int
	Speed      int
	Generation int
	Legendary  bool
}

type PokemonResponse struct {
	Name       string       `json:"name"`
	Cathegory  string       `json:"cathegory"`
	Generation int          `json:"generation"`
	Legendary  bool         `json:"legendary"`
	Score      PokemonScore `json:"score"`
}

type PokemonScore struct {
	HP      int `json:"hp"`
	Attack  int `json:"attack"`
	Defense int `json:"defense"`
	SpAtk   int `json:"spAtk"`
	SpDef   int `json:"spDef"`
	Speed   int `json:"speed"`
	Total   int `json:"total"`
}

const (
	COL_ID        = 0
	COL_NAME      = 1
	COL_TYPE      = 2
	COL_STYPE     = 3
	COL_TOTAL     = 4
	COL_HP        = 5
	COL_ATTACK    = 6
	COL_DEFENSE   = 7
	COL_SPATK     = 8
	COL_SPDEF     = 9
	COL_SPEED     = 10
	COL_GEN       = 11
	COL_LEGENDARY = 12
)

func BuildPokemonFromStore(line []string) Pokemon {
	id := parseInt(line[COL_ID])
	total := parseInt(line[COL_TOTAL])
	hp := parseInt(line[COL_HP])
	attack := parseInt(line[COL_ATTACK])
	defense := parseInt(line[COL_DEFENSE])
	spatk := parseInt(line[COL_SPATK])
	spdef := parseInt(line[COL_SPDEF])
	speed := parseInt(line[COL_SPEED])
	gen := parseInt(line[COL_GEN])
	legendary := parseBool(line[COL_LEGENDARY])
	return Pokemon{
		ID:         id,
		Name:       line[COL_NAME],
		Type:       line[COL_TYPE],
		Subtype:    line[COL_STYPE],
		Total:      total,
		HP:         hp,
		Attack:     attack,
		Defense:    defense,
		SpAtk:      spatk,
		SpDef:      spdef,
		Speed:      speed,
		Generation: gen,
		Legendary:  legendary,
	}
}

func parseInt(str string) int {
	parsed, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}
	return parsed
}

func parseBool(str string) bool {
	parsed, err := strconv.ParseBool(str)
	if err != nil {
		return false
	}
	return parsed
}
