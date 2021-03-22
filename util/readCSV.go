package util

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/joseantoniovz/academy-go-q12021/model"
)

const pathFile = "./booklist.csv"

func GetAll() ([]model.Book, *model.Error) {
	f, err := Open(pathFile)

	if err != nil {
		err := model.Error{
			Code:    http.StatusInternalServerError,
			Message: err,
		}
		return nil, &err
	}

	books, errorReading := ReadCsv(f)

	if errorReading != nil {
		errorReading := model.Error{
			Code:    http.StatusInternalServerError,
			Message: errorReading,
		}
		return nil, &errorReading
	}
	fmt.Println("books: ", books)
	return books, nil
}

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

		if line[4] != "" {
			price, err := strconv.ParseFloat(line[4], 64)
			//fmt.Println("error ", err)
			if err != nil {
				err := model.Error{
					Code:    http.StatusInternalServerError,
					Message: err.Error(),
				}
				return nil, &err
			}
			tempBook.Price = price
		}
		books = append(books, tempBook)
	}
	fmt.Println("Books", books)
	defer f.Close()
	return books, nil
}
