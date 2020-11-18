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

var (
	// DS Is the pokemon Datasource
	DS pokemon.DataSource
)

// CSV is the object used to read the csv files
type CSV struct {
	file *os.File
}

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
		r, e := reader.Read()

		if e == io.EOF {
			err = fmt.Errorf("Pokemon %s not found", name)
			return
		} else if e != nil {
			err = e
			return
		}

		if name == r[1] {
			p.Number, _ = strconv.Atoi(r[0])
			p.Name = r[1]
			p.Type1 = sql.NullString{String: r[2], Valid: r[2] != ""}
			p.Type2 = sql.NullString{String: r[3], Valid: r[3] != ""}
			p.HeightM, _ = strconv.ParseFloat(r[4], 64)
			p.WeightKg, _ = strconv.ParseFloat(r[5], 64)
			p.Male, _ = strconv.ParseFloat(r[6], 64)
			p.Female, _ = strconv.ParseFloat(r[7], 64)
			p.CaptRate, _ = strconv.ParseFloat(r[8], 64)
			p.HP, _ = strconv.Atoi(r[9])
			p.Attack, _ = strconv.Atoi(r[10])
			p.Defense, _ = strconv.Atoi(r[11])
			p.Special, _ = strconv.Atoi(r[12])
			p.Speed, _ = strconv.Atoi(r[13])
			return
		}
	}
}

// Close terminates the database connection
func (ds *CSV) Close() {
	ds.file.Close()
}
