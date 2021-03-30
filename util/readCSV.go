package util

import (
	"encoding/csv"
	"errors"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/fatih/structs"
	"github.com/joseantoniovz/academy-go-q12021/model"
)

func Open(path string) (*os.File, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, errors.New("There was an error opening the file")
	}
	return f, nil
}

func ReadCsv(f *os.File) ([]model.Book, *model.Error) {
	reader := csv.NewReader(f)
	_, err := reader.Read() // skip first line
	if err != nil {
		if err != io.EOF {
			log.Fatalln(err)
		}
	}
	var books []model.Book = nil
	for {
		line, err := reader.Read()

		if err == io.EOF {
			break
		}

		if err != nil {

			err := model.Error{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			}
			return nil, &err
		}

		tempBook := model.Book{
			Id:     line[0],
			Title:  line[1],
			Author: line[2],
			Format: line[3],
		}
		books = append(books, tempBook)
	}
	defer f.Close()
	return books, nil
}

func WriteInCSV(model model.Book, pathfile string) (*os.File, error) {

	s := make([]string, 0)
	f, err := os.OpenFile(pathfile, os.O_APPEND|os.O_WRONLY, 0600)

	if err != nil {
		return nil, errors.New("There was an error opening the file")
	}

	defer f.Close()

	writer := csv.NewWriter(f)
	for _, v := range structs.Values(model) {
		s = append(s, v.(string))
	}

	writer.Write(s)
	writer.Flush()
	return f, nil

}
