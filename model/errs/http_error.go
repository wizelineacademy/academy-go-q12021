package errs

import (
	"fmt"

	"github.com/grethelBello/academy-go-q12021/model"
)

type HttpError struct {
	model.HttpData
	StatusCode   int
	ErrorMessage string
}

func (e HttpError) Error() string {
	return fmt.Sprintf("Call to %v %v failed with status %v: %v", e.HttpData.Method, e.HttpData.Url, e.StatusCode, e.ErrorMessage)
}
