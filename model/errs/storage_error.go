package errs

import (
	"fmt"
)

type StorageError struct {
	TechnicalError error
}

func (e StorageError) Error() string {
	return fmt.Sprintf("Error accessing the storage")
}

func (e StorageError) Unwrap() error {
	return e.TechnicalError
}
