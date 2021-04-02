package repository

import "github.com/jesus-mata/academy-go-q12021/domain"

//go:generate mockgen -package mocks -destination $ROOTDIR/mocks/$GOPACKAGE/mock_$GOFILE . NewsArticleRepository
type NewsArticleRepository interface {
	FindAll() ([]*domain.NewsArticle, error)
	FindByID(id string) (*domain.NewsArticle, error)
	FetchCurrent() error
	GetIterator() (domain.NewsIterator, error)
}
