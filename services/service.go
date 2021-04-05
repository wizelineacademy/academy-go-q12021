package services

import (
	"encoding/csv"
	"os"
)

//Service struct with the parameters required in services
type Service struct {
	reader     *os.File
	writer     *csv.Writer
	pokemonAPI string
	pathFile   string
}

//New function to init the services
func New(reader *os.File, writer *csv.Writer, pokemonAPI string, pathFile string) (*Service, error) {
	return &Service{reader, writer, pokemonAPI, pathFile}, nil
}
