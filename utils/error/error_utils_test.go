package error

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	errorMessage = "error message"
)

func TestNewInternalServerError(t *testing.T) {
	err := NewInternalServerError(errorMessage)
	assert.NotNil(t, err)
	assert.EqualValues(t, 500, err.Code())
	assert.EqualValues(t, errorMessage, err.Message())
	assert.EqualValues(t, fmt.Sprintf("message: %s - code: %d", errorMessage, 500), err.Error())
}

func TestNewBadRequestError(t *testing.T) {
	err := NewBadRequestError(errorMessage)
	assert.NotNil(t, err)
	assert.EqualValues(t, 400, err.Code())
	assert.EqualValues(t, errorMessage, err.Message())
	assert.EqualValues(t, fmt.Sprintf("message: %s - code: %d", errorMessage, 400), err.Error())
}

func TestNewNotFoundError(t *testing.T) {
	err := NewNotFoundError(errorMessage)
	assert.NotNil(t, err)
	assert.EqualValues(t, 404, err.Code())
	assert.EqualValues(t, errorMessage, err.Message())
	assert.EqualValues(t, fmt.Sprintf("message: %s - code: %d", errorMessage, 404), err.Error())
}
