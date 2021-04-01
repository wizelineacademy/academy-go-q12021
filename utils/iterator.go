package utils

type Iterator interface {
	hasNext() bool
	getNext() [][]string
}
