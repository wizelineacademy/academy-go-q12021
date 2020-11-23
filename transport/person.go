package transport

import (
	"bufio"
	"encoding/csv"
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/pankecho/golang-bootcamp-2020/entity"
	"io"
	"net/http"
	"os"
	"time"
)

type PersonUseCase interface {
	UploadCSV(ctx echo.Context, data [][]string) ([]*entity.Person, error)
}

type Person struct {
	UseCase	PersonUseCase
}

func NewPerson(uc PersonUseCase) Person {
	return Person{
		UseCase: uc,
	}
}

func (p Person) UploadCSV(c echo.Context) error {
	req := c.Request()
	file, _, err := req.FormFile("file")
	if err != nil {
		return &echo.HTTPError{
			Code:     http.StatusBadRequest,
			Message:  "file not found",
		}
	}
	defer file.Close()

	tmpFileName := time.Now().String()
	f, err := os.OpenFile(tmpFileName, os.O_WRONLY | os.O_CREATE, 0666)
	if err != nil {
		return &echo.HTTPError{
			Code:     http.StatusInternalServerError,
			Message:  "unable open csv",
		}
	}
	defer f.Close()

	_, err = io.Copy(f, file)
	if err != nil {
		return &echo.HTTPError{
			Code:     http.StatusInternalServerError,
			Message:  "unable process csv",
		}
	}
	defer os.Remove(tmpFileName)

	records, err := parseCSV(tmpFileName)
	if err != nil {
		return &echo.HTTPError{
			Code:     http.StatusBadRequest,
			Message:  err.Error(),
		}
	}

	people, err := p.UseCase.UploadCSV(c, records)
	if err != nil {
		return &echo.HTTPError{
			Code:     http.StatusInternalServerError,
			Message:  err.Error(),
		}
	}

	return c.JSON(http.StatusCreated, people)
}

func parseCSV(filePath string) ([][]string, error) {
	var records [][]string

	f, err := os.Open(filePath)
	if err != nil {
		return records, errors.New("file not found")
	}

	r := csv.NewReader(bufio.NewReader(f))
	records, err = r.ReadAll()
	if err != nil {
		return records, errors.New("invalid csv format")
	}

	return records, nil
}