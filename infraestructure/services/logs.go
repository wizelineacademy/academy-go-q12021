package services

type Logger interface {
	Append(record string) error
	Get() ([]string, error)
}
