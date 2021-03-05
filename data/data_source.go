package data

type Source interface {
	init() ([][]string, error)
}