package data

import "github.com/grethelBello/academy-go-q12021/model"

// Source is an interface for the modules in charge of getting information from different sources (files, APIs, databases, etc.)
type Source interface {
	GetData(...*model.SourceConfig) (*model.Data, error)
	SetData(*model.Data) error
}
