package entity

import "time"

type Persona struct {
	ID        ID
	Nombre    string
	Apellido  string
	Edad      int
	Peso      float64
	Estatura  float64
	CreatedAt time.Time
}

func NewPersona(nombre, apellido string, edad int, peso, estatura float64) (*Persona, error) {
	p := &Persona{
		ID:        NewID(),
		Nombre:    nombre,
		Apellido:  apellido,
		Edad:      edad,
		Peso:      peso,
		Estatura:  estatura,
		CreatedAt: time.Now(),
	}
	if err := p.Validate(); err != nil {
		return nil, ErrInvalidEntity
	}
	return p, nil
}

func (p *Persona) Validate() error {
	if p.Nombre == "" || p.Apellido == "" || p.Edad <= 0 || p.Peso <= 0 || p.Estatura < 0 {
		return ErrInvalidEntity
	}
	return nil
}
