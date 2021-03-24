package query

import (
	"github.com/maestre3d/academy-go-q12021/internal/aggregate"
	"github.com/maestre3d/academy-go-q12021/internal/valueobject"
)

// MovieResponse Movie's DTO used for representation layer(s)
type MovieResponse struct {
	MovieID     string   `json:"movie_id"`
	DisplayName string   `json:"title"`
	Directors   []string `json:"directors"`
	ReleaseYear int      `json:"release_year"`
}

// MoviesResponse A list of movies as DTO used for representation layer(s)
type MoviesResponse struct {
	TotalItems int             `json:"total_items"`
	Movies     []MovieResponse `json:"movies"`
	NextPage   string          `json:"next_page"`
}

func marshalMovieResponse(m *aggregate.Movie) MovieResponse {
	return MovieResponse{
		MovieID:     string(m.ID),
		DisplayName: string(m.DisplayName),
		Directors:   valueobject.MarshalDirectorsPrimitive(m.Directors...),
		ReleaseYear: int(m.ReleaseYear),
	}
}

func marshalMoviesResponse(nextPage string, movies ...*aggregate.Movie) MoviesResponse {
	moviesResp := make([]MovieResponse, 0)
	for _, m := range movies {
		moviesResp = append(moviesResp, marshalMovieResponse(m))
	}

	return MoviesResponse{
		Movies:     moviesResp,
		TotalItems: len(moviesResp),
		NextPage:   nextPage,
	}
}
