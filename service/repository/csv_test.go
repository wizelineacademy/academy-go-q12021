package repository

import (
	"testing"

	"github.com/etyberick/golang-bootcamp-2020/entity"
)

func TestRepository(t *testing.T) {
	qr := NewQuoteRepository("test.csv")
	q := &entity.Quote{
		ID:     "1337",
		Text:   "Test succesfull",
		Author: "Etyberick",
		Genre:  "Test",
	}
	qr.Write(*q)
	qs, err := qr.ReadAll()
	if err != nil {
		t.Fatalf(err.Error())
	}
	if len(qs) <= 0 {
		t.Fatalf("unexpected result size: %d", len(qs))
	}
}
