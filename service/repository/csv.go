package repository

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/etyberick/golang-bootcamp-2020/entity"
)

const ownerRWX = 755

type quoteRepository struct {
	file   os.File
	quotes []entity.Quote
}

// ErrEmptyDatabase message response
const ErrEmptyDatabase = "database is empty"

// QuoteRepository manages all persistency management operations
type QuoteRepository interface {
	ReadAll() ([]entity.Quote, error)
	Write(entity.Quote) error
}

// NewQuoteRepository will create an entity that represents the entity.QuoteStorage interface
func NewQuoteRepository(filename string) QuoteRepository {
	// Create file when it doesn't exist or load it when it exists
	qr := &quoteRepository{}
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		_, err = os.Create(filename)
		if err != nil {
			log.Printf("error creating %s", filename)
			log.Fatalf("%s", err)
		}
	} else {
		qr.loadCache(filename)
	}

	// Open the file and it to the new object
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, ownerRWX)
	if err != nil {
		log.Printf("error opening %s", filename)
		log.Fatalf("%s", err)
	}
	qr.file = *f
	return qr
}

// ReadAll items from data source
func (qr *quoteRepository) ReadAll() ([]entity.Quote, error) {
	if len(qr.quotes) == 0 {
		return nil, fmt.Errorf(ErrEmptyDatabase)
	}
	return qr.quotes, nil
}

// Write an item into data source
func (qr *quoteRepository) Write(quote entity.Quote) error {
	// Write into the CSV file
	csvw := csv.NewWriter(&qr.file)
	if err := csvw.Write(quote.Slice()); err != nil {
		log.Println(err)
		return err
	}
	csvw.Flush()

	// Append it into the quotes slice
	qr.quotes = append(qr.quotes, quote)

	return nil
}

func (qr *quoteRepository) loadCache(filename string) {
	//Open the file
	f, err := os.Open(filename)
	if err != nil {
		log.Printf("Warning %s, skipping cache loading", err.Error())
	}
	defer f.Close()

	//Read it
	csvr := csv.NewReader(bufio.NewReader(f))
	for {
		r, err := csvr.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("error reading line: %s", err)
		}
		if l := len(r); l < 4 {
			log.Printf("incomplete line, has %d instead of 4", l)
			continue
		}

		q := entity.Quote{
			ID:     r[0],
			Text:   r[1],
			Author: r[2],
			Genre:  r[3],
		}
		qr.quotes = append(qr.quotes, q)
	}

}
