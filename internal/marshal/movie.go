package marshal

import (
	"strings"

	"github.com/maestre3d/academy-go-q12021/internal/domain"
	"github.com/maestre3d/academy-go-q12021/internal/valueobject"
)

// ErrCannotParseMovie the given movie could not get parsed
var ErrCannotParseMovie = domain.NewInfrastructure("cannot parse movie")

func unmarshalMovieDirectors(delimeter, v string) []valueobject.DisplayName {
	directors := make([]valueobject.DisplayName, 0)
	splitDirectors := strings.Split(v, delimeter)
	for _, d := range splitDirectors {
		directors = append(directors, valueobject.DisplayName(d))
	}

	return directors
}
