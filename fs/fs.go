package fs

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/adantop/golang-bootcamp-2020/pokemon"
)

// DS Is the pokemon Datasource
var DS pokemon.DataSource

type (
	// CSV is the object used to read the csv files
	CSV    struct{ file *os.File }
	record []string
)

// UseCSV generates the CSV object to be used
func UseCSV(filename string) {
	file, err := os.Open(filename)

	if err != nil {
		log.Fatalln(err)
	}

	DS = &CSV{file}
}

// GetPokemonByName Get pokemon by name
func (ds *CSV) GetPokemonByName(name string) (p pokemon.Pokemon, err error) {
	reader := csv.NewReader(ds.file)

	// skip header
	reader.Read()

	for {
		var r record
		r, e := reader.Read()

		if e == io.EOF {
			err = fmt.Errorf("Pokemon %s not found", name)
			return
		}

		if e != nil {
			err = e
			return
		}

		if name == r[1] {
			return r.producePokemon()
		}
	}
}

// Close terminates the database connection
func (ds *CSV) Close() {
	ds.file.Close()
}

func (r *record) producePokemon() (p pokemon.Pokemon, err error) {

	if p.Number, err = strconv.Atoi((*r)[0]); err != nil {
		err = fmt.Errorf("Could not parse int for pokemon Number, column 1")
		return
	}

	p.Name = (*r)[1]
	p.Type1 = sql.NullString{String: (*r)[2], Valid: (*r)[2] != ""}
	p.Type2 = sql.NullString{String: (*r)[3], Valid: (*r)[3] != ""}

	if p.HeightM, err = strconv.ParseFloat((*r)[4], 64); err != nil {
		err = fmt.Errorf("Could not parse float for pokemon Height, column 5")
		return
	}

	if p.WeightKg, err = strconv.ParseFloat((*r)[5], 64); err != nil {
		err = fmt.Errorf("Could not parse float for pokemon Weight, column 6")
		return
	}

	if p.Male, err = strconv.ParseFloat((*r)[6], 64); err != nil {
		err = fmt.Errorf("Could not parse float for pokemon Male ratio, column 7")
		return
	}

	if p.Female, err = strconv.ParseFloat((*r)[7], 64); err != nil {
		err = fmt.Errorf("Could not parse float for pokemon Female ratio, column 8")
		return
	}

	if p.CaptRate, err = strconv.ParseFloat((*r)[8], 64); err != nil {
		err = fmt.Errorf("Could not parse float for pokemon Capture rate, column 6")
		return
	}

	if p.HP, err = strconv.Atoi((*r)[9]); err != nil {
		err = fmt.Errorf("Could not parse int for pokemon HP, column 10")
		return
	}

	if p.Attack, err = strconv.Atoi((*r)[10]); err != nil {
		err = fmt.Errorf("Could not parse int for pokemon Attack, column 11")
		return
	}

	if p.Defense, err = strconv.Atoi((*r)[11]); err != nil {
		err = fmt.Errorf("Could not parse int for pokemon Defense, column 12")
		return
	}

	if p.Special, err = strconv.Atoi((*r)[12]); err != nil {
		err = fmt.Errorf("Could not parse int for pokemon Special, column 13")
		return
	}

	if p.Speed, err = strconv.Atoi((*r)[13]); err != nil {
		err = fmt.Errorf("Could not parse int for pokemon Speed, column 14")
		return
	}
	return
}
