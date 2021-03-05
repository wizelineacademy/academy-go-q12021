package repository

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/jesus-mata/academy-go-q12021/domain"
	"github.com/jesus-mata/academy-go-q12021/infrastructure"
	"github.com/jesus-mata/academy-go-q12021/usecase/repository"
	"github.com/jesus-mata/academy-go-q12021/utils"
)

type newsRepository struct {
	csvReader *infrastructure.CsvReader
	logger    *log.Logger
}

func NewNewsArticleRepository(csv *infrastructure.CsvReader, logger *log.Logger) repository.NewsArticleRepository {
	return &newsRepository{csv, logger}
}

func (r *newsRepository) FindByID(id int) (*domain.NewsArticle, error) {
	r.logger.Println("Finding News Article by ID", id)
	records, err := r.csvReader.GetAllLines()
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
	records, err := r.csvReader.GetAllLines()
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

func mapFromCSVRecord(record []string) (*domain.NewsArticle, error) {
	if len(record) < 4 {
		return nil, errors.New("CSV file has not valid News data")
	}

	id, err := strconv.Atoi(record[0])
	if err != nil {
		return nil, fmt.Errorf("CSV has invalid data. The ID '%v' is not valid.", record[0])
	}
	author := record[1]
	category := record[2]
	publishedDate, err := time.Parse(utils.LayoutDateTimeISO, record[3])
	if err != nil {
		return nil, fmt.Errorf("CSV has invalid data. The Published Date '%v' is not a valid date.", record[3])
	}

	title := record[4]

	return domain.CreateNewsArticle(id, publishedDate, title, category, author)
}
