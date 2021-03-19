package errs

import (
	"fmt"
)

// StorageError is a wrapper for any error related to the datasources
type StorageError struct {
	TechnicalError error
}

func (e StorageError) Error() string {
	return fmt.Sprintf("Error accessing the storage: %v", e.TechnicalError.Error())
}

func (e StorageError) Unwrap() error {
	return e.TechnicalError
}
