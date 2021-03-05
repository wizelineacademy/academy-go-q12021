package interactors

import (
	"github.com/jesus-mata/academy-go-q12021/domain"
	"github.com/jesus-mata/academy-go-q12021/usecase/repository"
)

type NewsArticlesInteractor interface {
	GetAll() ([]*domain.NewsArticle, error)
	GetByID(id int) (*domain.NewsArticle, error)
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

func (s *newsArticlesInteractor) GetByID(id int) (*domain.NewsArticle, error) {
	return s.newsRepository.FindByID(id)
}
