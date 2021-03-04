package application

import (
	"context"
	"testing"

	"github.com/maestre3d/academy-go-q12021/internal/aggregate"
	"github.com/maestre3d/academy-go-q12021/internal/domain"
	"github.com/maestre3d/academy-go-q12021/internal/persistence"
	"github.com/maestre3d/academy-go-q12021/internal/repository"
	"github.com/maestre3d/academy-go-q12021/internal/valueobject"
	"github.com/stretchr/testify/assert"
)

func TestMovieGetByID(t *testing.T) {
	movieApp := NewMovie(persistence.NewMovieInMemory(), nil, nil)
	movie, err := movieApp.GetByID(context.Background(), valueobject.MovieID("1"))
	assert.NotNil(t, err)
	errDomain := err.(domain.Error)
	assert.True(t, errDomain.IsNotFound())
	assert.Nil(t, movie)

	movieStub := aggregate.NewEmptyMovie()
	movieStub.ID = valueobject.MovieID("1")
	movieStub.DisplayName = valueobject.DisplayName("Foo Movie")
	_ = movieApp.repo.Save(context.Background(), *movieStub)

	movie, err = movieApp.GetByID(context.Background(), movieStub.ID)
	assert.Nil(t, err)
	assert.EqualValues(t, movieStub, movie)
}

func TestMovieList(t *testing.T) {
	repo := persistence.NewMovieInMemory()
	// NewCriteria sets 100 to limit by default if empty value was given
	criteria := *repository.NewCriteria(0, "")
	movieApp := NewMovie(repo, nil, nil)
	movies, _, err := movieApp.List(context.Background(), criteria)
	assert.NotNil(t, err)
	errDomain := err.(domain.Error)
	assert.True(t, errDomain.IsNotFound())
	assert.Equal(t, 0, len(movies))

	movieStub := aggregate.NewEmptyMovie()
	movieStub.ID = valueobject.MovieID("1")
	movieStub.DisplayName = valueobject.DisplayName("Foo Movie")
	_ = repo.Save(context.Background(), *movieStub)

	movies, _, err = movieApp.List(context.Background(), criteria)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(movies))
}
