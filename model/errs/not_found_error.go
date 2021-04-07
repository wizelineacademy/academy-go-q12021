package errs

import (
	"fmt"
)

type NotFoundError struct {
	Id             int
	Datatype       string
	TechnicalError error
}

func (e NotFoundError) Error() string {
	return fmt.Sprintf("The %v with %v ID was not found", e.Datatype, e.Id)
}

func (e NotFoundError) Unwrap() error {
	return e.TechnicalError
}
