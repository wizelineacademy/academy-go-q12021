package interactors

import (
	"github.com/jesus-mata/academy-go-q12021/application/repository"
	"github.com/jesus-mata/academy-go-q12021/domain"
)

type NewsArticlesInteractor interface {
	GetAll() ([]*domain.NewsArticle, error)
	GetByID(id string) (*domain.NewsArticle, error)
	FetchAll() error
}

type newsArticlesInteractor struct {
	newsRepository repository.NewsArticleRepository
}

func NewNewsArticlesInteractor(newsRepository repository.NewsArticleRepository) NewsArticlesInteractor {
	return &newsArticlesInteractor{newsRepository}
}

func (s *newsArticlesInteractor) GetAll() ([]*domain.NewsArticle, error) {
	return s.newsRepository.FindAll()
}

func (s *newsArticlesInteractor) GetByID(id string) (*domain.NewsArticle, error) {
	return s.newsRepository.FindByID(id)
}

func (s *newsArticlesInteractor) FetchAll() error {
	return s.newsRepository.FetchCurrent()
}

func (s *newsArticlesInteractor) FetchAllConcurrenlty() error {
	it := s.newsRepository.FetchCurrent()
	it.
	return error
}
