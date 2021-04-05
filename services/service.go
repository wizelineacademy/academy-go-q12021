package services

import (
	"encoding/csv"
	"os"
)

type Service struct {
	reader     *os.File
	writer     *csv.Writer
	pokemonAPI string
	pathFile   string
}

func New(reader *os.File, writer *csv.Writer, pokemonAPI string, pathFile string) (*Service, error) {
	return &Service{reader, writer, pokemonAPI, pathFile}, nil
}
