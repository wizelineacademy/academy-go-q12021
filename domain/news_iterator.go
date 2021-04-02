package domain

type NewsIterator interface {
	HasNext() (bool, error)
	GetNext() *NewsArticle
}
