package errs

import (
	"fmt"

	"github.com/wizelineacademy/academy-go/model"
)

type HttpError struct {
	model.HttpData
	StatusCode   int
	ErrorMessage string
}

func (e HttpError) Error() string {
	return fmt.Sprintf("Call to %v %v failed with status %v: %v", e.HttpData.Method, e.HttpData.Url, e.StatusCode, e.ErrorMessage)
}
