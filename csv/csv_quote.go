package csv

import (
	"fmt"
	"log"
	"os"

	"github.com/etyberick/golang-bootcamp-2020/entity"
)

type csvQuoteRepository struct {
	Filename string
}

//NewCsvQuoteRepository will create an entity that represents the entity.QuoteStorage interface
func NewCsvQuoteRepository(filename string) entity.QuoteStorage {
	return &csvQuoteRepository{filename}
}

func (csvQuoteRepo *csvQuoteRepository) ReadAll() ([]entity.Quote, error) {
	panic("ReadAll not implemented") // TODO: Implement
}

func (csvQuoteRepo *csvQuoteRepository) Write(quote *entity.Quote) error {
	f, err := os.OpenFile(csvQuoteRepo.Filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 755)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	csvRow := fmt.Sprintf("%q, %q, %q, %q\n", quote.ID, quote.QuoteText, quote.QuoteAuthor, quote.QuoteGenre)
	if _, err := f.WriteString(csvRow); err != nil {
		log.Println(err)
	}
	return nil
}
