package validator

import (
	"github.com/go-playground/validator/v10"
)

type customValidator struct {
	validator *validator.Validate
}

type CustomValidator interface {
	Validate(i interface{}) error
}

func NewCustomValidator() CustomValidator {
	return &customValidator{validator.New()}
}

func (cv *customValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return err
	}
	return nil
}
