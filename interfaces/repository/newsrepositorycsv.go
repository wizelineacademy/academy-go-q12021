package repository

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/jesus-mata/academy-go-q12021/application/repository"
	"github.com/jesus-mata/academy-go-q12021/domain"
	"github.com/jesus-mata/academy-go-q12021/infrastructure"
	"github.com/jesus-mata/academy-go-q12021/infrastructure/newsapi"
	"github.com/jesus-mata/academy-go-q12021/utils"
)

type newsRepository struct {
	csvSource infrastructure.CsvSource
	newsApi   newsapi.NewsApiClient
	logger    *log.Logger
}

func NewNewsArticleRepository(csv infrastructure.CsvSource, newsApi newsapi.NewsApiClient, logger *log.Logger) repository.NewsArticleRepository {
	return &newsRepository{csv, newsApi, logger}
}

func (r *newsRepository) FindByID(id string) (*domain.NewsArticle, error) {
	r.logger.Println("Finding News Article by ID", id)
	records, err := r.csvSource.GetAllLines()
	if err != nil {
		return nil, err
	}

	for _, record := range records {
		newsArticle, err := mapFromCSVRecord(record)
		if err != nil {
			return nil, err
		}

		if newsArticle.Id == id {
			return newsArticle, nil
		}
	}

	return nil, errors.New(fmt.Sprintf("News Article with ID %v does not exist.", id))
}

func (r *newsRepository) FindAll() ([]*domain.NewsArticle, error) {
	r.logger.Println("Retriving all News Articles")
	records, err := r.csvSource.GetAllLines()
	if err != nil {
		return nil, err
	}

	newsArticles := make([]*domain.NewsArticle, 0, 5)

	for _, record := range records {
		newsArticle, err := mapFromCSVRecord(record)
		if err != nil {
			return nil, err
		}

		newsArticles = append(newsArticles, newsArticle)
	}

	return newsArticles, nil
}

func (r *newsRepository) FetchCurrent() error {
	r.logger.Println("Fetching all News Articles from API")
	newsItems, err := r.newsApi.GetCurrentNews()
	if err != nil {
		return err
	}
	r.logger.Printf("News Found %v \n", len(newsItems))
	err = r.csvSource.WriteLines(newsItems)
	if err != nil {
		return err
	}
	return nil
}

func mapFromCSVRecord(record []string) (*domain.NewsArticle, error) {
	if len(record) < 4 {
		return nil, errors.New("CSV file has not valid News data")
	}

	id := record[0]
	title := record[1]
	description := record[2]
	url := record[3]
	author := record[4]
	image := record[5]
	language := record[6]
	category := record[7]
	publishedDate, err := time.Parse(utils.LayoutDateTimeIDOWithTZ, record[8])
	if err != nil {
		return nil, fmt.Errorf("CSV has invalid data. The Published Date '%v' is not a valid date.", record[3])
	}

	return domain.CreateNewsArticle(id, title, description, url, author, image, language, category, publishedDate)
}
