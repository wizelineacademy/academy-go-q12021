package errors

import (
    "fmt"
)

type NoDataError struct {
    ErrMsg   error
}

func (e NoDataError) Error() string {
    return fmt.Sprintf("Data unavailable: %v", e.ErrMsg)
}

func (e NoDataError) Unwrap() error {
    return e.ErrMsg
}
