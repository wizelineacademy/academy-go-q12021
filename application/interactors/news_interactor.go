package interactors

import (
	"fmt"
	"sync"

	"github.com/jesus-mata/academy-go-q12021/application/repository"
	"github.com/jesus-mata/academy-go-q12021/domain"
	"github.com/jesus-mata/academy-go-q12021/utils/worker"
)

type NewsArticlesInteractor interface {
	GetAll() ([]*domain.NewsArticle, error)
	GetByID(id string) (*domain.NewsArticle, error)
	FetchAll() error
	FindAllByCategoryConcurrenlty(category string, limit, itemsPerWorker int) ([]*domain.NewsArticle, error)
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

func (s *newsArticlesInteractor) FindAllByCategoryConcurrenlty(category string, limit, itemsPerWorker int) ([]*domain.NewsArticle, error) {
	articles := make([]*domain.NewsArticle, 0, 5)

	it, err := s.newsRepository.GetIterator()
	if err != nil {
		return articles, err
	}

	hasNext, err := it.HasNext()
	if err != nil {
		return articles, err
	}

	var wg sync.WaitGroup

	results := make(chan *domain.NewsArticle, limit)

	workerPool := worker.NewWorkerPool(limit, itemsPerWorker, results, &wg)
	workerPool.Start()

	go func() error {
		for hasNext {

			news := it.GetNext()
			fmt.Println("news: ", news)
			job := worker.NewNewsJobFilter(news, category)
			workerPool.AddJob(job)

			hasNext, err = it.HasNext()
			if err != nil {
				return err
			}

		}
		workerPool.ShutDown()
		return nil
	}()

	for news := range results {
		fmt.Println("Articles range: ", articles)
		articles = append(articles, news)
		if len(articles) == limit {
			workerPool.Stop()
			break
		}
	}
	workerPool.Stop()
	fmt.Println("Articles: ", articles)

	return articles, nil
}
