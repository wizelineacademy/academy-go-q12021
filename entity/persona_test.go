package entity_test

import (
	"testing"

	"github.com/fdsmora/golang-bootcamp-2020/entity"
	"github.com/stretchr/testify/assert"
)

func TestNewPersona(t *testing.T) {
	p, err := entity.NewPersona("Cuco", "Sanchez", 54, 59, 1.60)
	assert.Nil(t, err)
	assert.Equal(t, p.Nombre, "Cuco")
	assert.NotNil(t, p.ID)
}

func TestPersonaValidate(t *testing.T) {
	type test struct {
		nombre   string
		apellido string
		edad     int
		peso     float64
		estatura float64
		want     error
	}

	tests := []test{
		{
			nombre:   "Fausto",
			apellido: "Salazar",
			edad:     33,
			peso:     75.0,
			estatura: 1.75,
			want:     nil,
		},
		{
			nombre:   "Fausto",
			apellido: "Salazar",
			edad:     0,
			peso:     75.0,
			estatura: 1.75,
			want:     entity.ErrInvalidEntity,
		},
		{
			nombre:   "",
			apellido: "Salazar",
			edad:     10,
			peso:     75.0,
			estatura: 1.75,
			want:     entity.ErrInvalidEntity,
		},
		{
			nombre:   "Fausto",
			apellido: "Salazar",
			edad:     0,
			peso:     -75.0,
			estatura: 1.75,
			want:     entity.ErrInvalidEntity,
		},
	}
	for _, tc := range tests {
		_, err := entity.NewPersona(tc.nombre, tc.apellido, tc.edad, tc.peso, tc.estatura)
		assert.Equal(t, err, tc.want)
	}
}
