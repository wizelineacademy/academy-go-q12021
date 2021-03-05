package domain

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

type NewsArticle struct {
	Id            int       `json:"id"`
	Author        string    `json:"author"`
	Category      string    `json:"category"`
	PublishedDate time.Time `json:"publishedDate"`
	Title         string    `json:"title"`
}

func CreateNewsArticle(id int, publishedDate time.Time, title string, category string, author string) (*NewsArticle, error) {

	if id <= 0 {
		return nil, errors.New(fmt.Sprintf("News ID %v is not a valid ID.", id))
	}

	if len(strings.TrimSpace(title)) == 0 {
		return nil, errors.New("News Title for the article is mandatory.")
	}

	return &NewsArticle{Id: id, PublishedDate: publishedDate, Title: title, Category: category, Author: author}, nil
}
