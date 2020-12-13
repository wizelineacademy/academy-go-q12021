package router

import (
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestRoute1(t *testing.T) {
	r := mux.NewRouter()
	path, err := r.GetRoute("readcsv").GetPathTemplate()
	assert.Nil(t, err)
	assert.Equal(t, "/api/readcsv", path)
}
