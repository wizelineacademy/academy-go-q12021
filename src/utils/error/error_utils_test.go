package error

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

const errorMessage = "error message"

func TestNewInternalServerError(t *testing.T) {
	err := NewInternalServerError(errorMessage)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Code())
	assert.EqualValues(t, errorMessage, err.Message())
	assert.EqualValues(t, fmt.Sprintf("message: %s - code: %d", errorMessage, http.StatusInternalServerError), err.Error())
}

func TestNewBadRequestError(t *testing.T) {
	err := NewBadRequestError(errorMessage)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusBadRequest, err.Code())
	assert.EqualValues(t, errorMessage, err.Message())
	assert.EqualValues(t, fmt.Sprintf("message: %s - code: %d", errorMessage, http.StatusBadRequest), err.Error())
}

func TestNewNotFoundError(t *testing.T) {
	err := NewNotFoundError(errorMessage)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.Code())
	assert.EqualValues(t, errorMessage, err.Message())
	assert.EqualValues(t, fmt.Sprintf("message: %s - code: %d", errorMessage, http.StatusNotFound), err.Error())
}
