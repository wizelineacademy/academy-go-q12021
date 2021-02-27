package core

type Error interface {
	Error() string
}

type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return e.s
}

func NewError(text string) error {
	return &errorString{text}
}
