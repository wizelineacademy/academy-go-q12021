package worker

import "github.com/jesus-mata/academy-go-q12021/domain"

type Job interface {
	Apply() bool
	GetData() *domain.NewsArticle
}

type NewsJobFilter struct {
	data     *domain.NewsArticle
	category string
}

func NewNewsJobFilter(data *domain.NewsArticle, category string) Job {
	return &NewsJobFilter{data: data, category: category}
}

func (n NewsJobFilter) Apply() bool {
	if n.data.Category == n.category {
		return true
	}
	return false
}

func (n NewsJobFilter) GetData() *domain.NewsArticle {
	return n.data
}
