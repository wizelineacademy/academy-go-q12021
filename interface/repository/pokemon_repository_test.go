package repository_test

import (
	"testing"

	"github.com/Topi99/academy-go-q12021/interface/repository"
	uc_repository "github.com/Topi99/academy-go-q12021/usecase/repository"
	"github.com/stretchr/testify/assert"
)

func TestFindOne(t *testing.T) {
	type test struct {
		id   uint
		want error
	}

	tests := []test{
		{
			id:   1,
			want: nil,
		},
		{
			id:   500,
			want: uc_repository.ErrNotFound,
		},
	}

	for _, tp := range tests {
		r := repository.NewPokemonRepository()
		_, err := r.FindOne(tp.id)

		assert.Equal(t, err, tp.want)
	}
}
