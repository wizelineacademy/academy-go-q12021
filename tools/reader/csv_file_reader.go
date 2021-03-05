package csv_file_reader

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/ToteEmmanuel/academy-go-q12021/domain/model"
)

type CsvPokeStorage struct {
	pokeMap map[int32]*model.Pokemon
}

func NewCsvStorage(file string) (*CsvPokeStorage, error) {
	source, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	reader := csv.NewReader(source)
	pokeMap := make(map[int32]*model.Pokemon)
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
			continue
		}
		if len(record) != 5 {
			fmt.Errorf("Incorrect amount of fields(5) in :", record)
		}
		var id int32
		if val, err := strconv.Atoi(record[0]); err != nil {
			fmt.Errorf("Error  parsing line ", record)
			continue
		} else {
			id = int32(val)
		}
		pokeMap[id] = &model.Pokemon{
			Id:         id,
			Species:    record[1],
			Sprite:     record[2],
			Types:      strings.Split(record[3], ","),
			FlavorText: record[4],
		}
	}
	return &CsvPokeStorage{pokeMap}, nil
}

func (c *CsvPokeStorage) FindById(id int32) *model.Pokemon {
	return c.pokeMap[id]
}

func (c *CsvPokeStorage) FindAll() []*model.Pokemon {
	allPokes := make([]*model.Pokemon, 0, len(c.pokeMap))
	for _, v := range c.pokeMap {
		allPokes = append(allPokes, v)
	}
	return allPokes
}
