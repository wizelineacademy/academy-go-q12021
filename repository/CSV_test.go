package CSV_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadPersonaFromCSV(t *testing.T) {
	type test struct {
		ID,
		nombre,
		apellido,
		edad,
		peso,
		estatura string
	}
	tests := []test{
		{
			ID:       "id",
			nombre:   "nombre",
			apellido: "apellido",
			edad:     "edad",
			peso:     "peso",
			estatura: "estatura",
		},
		{
			ID:       "1",
			nombre:   "fausto",
			apellido: "salazar",
			edad:     "33",
			peso:     "75",
			estatura: "1.75",
		},
	}
	rows := CSV.ReadFile()
	for i, want := range tests {
		assert.Equal(t, want, rows[i])
	}
}
