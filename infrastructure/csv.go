package infrastructure

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
)

type CsvReader struct {
	file   string
	logger *log.Logger
}

func NewCsvReader(file string, logger *log.Logger) *CsvReader {

	return &CsvReader{file, logger}
}

func (c *CsvReader) PrintAllLines() error {

	csvfile, err := os.Open(c.file)
	if err != nil {
		return err
	}

	reader := csv.NewReader(csvfile)
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			c.logger.Fatal(err)
			return err
		}
		c.logger.Println("CSV Record: ", record)
	}
	return nil
}

func (c *CsvReader) GetAllLines() ([][]string, error) {
	rows := make([][]string, 0, 5)

	c.logger.Panicln("Reading all lines of CSV file", c.file)

	csvfile, err := os.Open(c.file)
	if err != nil {
		return nil, err
	}
	reader := csv.NewReader(csvfile)

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err, ok := err.(*csv.ParseError); ok {
			return nil, errors.New(fmt.Sprintf("Cannot parse CSV: %s", err.Error()))
		}
		if err != nil {
			return nil, err
		}
		rows = append(rows, record)
	}
	return rows, nil
}
