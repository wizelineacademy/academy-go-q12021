package infrastructure

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/jesus-mata/academy-go-q12021/infrastructure/dto"
)

type CsvSource struct {
	file   string
	logger *log.Logger
}

func NewCsvSource(file string, logger *log.Logger) *CsvSource {

	return &CsvSource{file, logger}
}

func (c *CsvSource) PrintAllLines() error {

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

func (c *CsvSource) GetAllLines() ([][]string, error) {
	rows := make([][]string, 0, 5)

	c.logger.Println("Reading all lines of CSV file", c.file)

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

func (c *CsvSource) WriteLines(newsItems []dto.NewItem) error {
	f, err := os.OpenFile(c.file, os.O_TRUNC|os.O_WRONLY, os.ModeAppend)
	defer f.Close()

	if err != nil {
		return err
	}

	w := csv.NewWriter(f)
	defer w.Flush()

	for _, newItem := range newsItems {
		category := ""
		if len(newItem.Category) > 0 {
			category = newItem.Category[0]
		}
		record := []string{newItem.Id, newItem.Title, newItem.Description, newItem.Url, newItem.Author, newItem.Image, newItem.Language, category, newItem.Published}
		if err := w.Write(record); err != nil {
			fmt.Println(err)
			return err
		}
	}

	return nil
}
