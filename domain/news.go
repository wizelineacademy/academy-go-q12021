package domain

import (
	"errors"
	"strings"
	"time"
)

type NewsArticle struct {
	Id            string    `json:"id"`
	Title         string    `json:"title"`
	Description   string    `json:"description"`
	Url           string    `json:"url"`
	Author        string    `json:"author"`
	Image         string    `json:"image"`
	Language      string    `json:"language"`
	Category      string    `json:"category"`
	PublishedDate time.Time `json:"publishedDate"`
}

func CreateNewsArticle(id, title, description, url, author, image, language, category string, publishedDate time.Time) (*NewsArticle, error) {

	if len(strings.TrimSpace(title)) == 0 {
		return nil, errors.New("News Title for the article is mandatory.")
	}

	return &NewsArticle{Id: id, PublishedDate: publishedDate, Title: title, Category: category, Author: author, Description: description, Image: image, Url: url, Language: language}, nil
}
