package poke

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

// PokemonList is the Pokemon dictionary from the CSV
type PokemonList map[string][]Pokemon

// Pokemon represents a Pokemon from the CSV file
type Pokemon struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	Type1          string `json:"type1"`
	Type2          string `json:"type2"`
	Total          string `json:"total"`
	HP             string `json:"hp"`
	Attack         string `json:"attack"`
	Defense        string `json:"defense"`
	SpecialAttack  string `json:"special_attack"`
	SpecialDefense string `json:"special_defense"`
	Speed          string `json:"speed"`
	Generation     string `json:"generation"`
	Legendary      string `json:"legendary"`
}

// Service implements the logic to get the required pokemon from a CSV
type Service struct {
	file        string
	PokemonList PokemonList
	rawList     [][]string
}

// NewPokeService creates a Service using the specified CSV file
func NewPokeService(file string) (*Service, error) {
	ps := Service{
		file: file,
	}
	err := ps.readCSV()
	if err != nil {
		return nil, err
	}
	ps.PokemonList = ps.assignPoke(ps.rawList)
	return &ps, nil
}

// GetPokeByID search for the pokemon id in the loaded Poke Dictionary
func (ps *Service) GetPokeByID(ID int) ([]Pokemon, error) {
	pid := strconv.Itoa(ID)
	return ps.PokemonList[pid], nil
}

func (ps *Service) assignPoke(pokes [][]string) PokemonList {
	pl := make(PokemonList)

	for i, p := range pokes {
		if i != 0 {
			pokemon := Pokemon{
				ID:             p[0],
				Name:           p[1],
				Type1:          p[2],
				Type2:          p[3],
				Total:          p[4],
				HP:             p[5],
				Attack:         p[6],
				Defense:        p[7],
				SpecialAttack:  p[8],
				SpecialDefense: p[9],
				Speed:          p[10],
				Generation:     p[11],
				Legendary:      p[12],
			}
			pt := pl[pokemon.ID]
			pt = append(pt, pokemon)
			pl[pokemon.ID] = pt
		}
	}

	return pl
}

func (ps *Service) readCSV() error {
	f, err := os.Open(ps.file)
	if err != nil {
		return fmt.Errorf("unable to read csv file: %v", err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		return fmt.Errorf("unable to parse csv: %v", err)
	}
	ps.rawList = records
	return nil
}
