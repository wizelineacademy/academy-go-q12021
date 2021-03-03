package poke

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
)

type PokeList map[string][]Pokemon

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

type PokeService struct {
	file        string
	PokemonList PokeList
	rawList     [][]string
}

func NewPokeService(file string) *PokeService {
	ps := PokeService{
		file: file,
	}
	ps.readCSV()
	ps.PokemonList = ps.assignPoke(ps.rawList)
	return &ps
}

func (ps *PokeService) GetPokeByID(ID int) ([]Pokemon, error) {
	pid := strconv.Itoa(ID)
	return ps.PokemonList[pid], nil
}

func (ps *PokeService) assignPoke(pokes [][]string) PokeList {
	pl := make(PokeList)

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

func (ps *PokeService) readCSV() {
	f, err := os.Open(ps.file)
	if err != nil {
		log.Fatalln("unable to read csv", err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatalln("unable to parse csv: ", err)
	}
	ps.rawList = records
}
