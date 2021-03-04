package command

import (
	"context"

	"github.com/maestre3d/academy-go-q12021/internal/application"
	"github.com/maestre3d/academy-go-q12021/internal/valueobject"
)

// SyncMovie requests a movie fetch and update from an external source
type SyncMovie struct {
	ID     string `json:"movie_id"`
	IMDbID string `json:"imdb_id"`
}

// HandleSyncMovie executes a SyncMovie command
func HandleSyncMovie(ctx context.Context, app *application.Movie, cmd SyncMovie) error {
	id, err := valueobject.NewMovieID(cmd.ID)
	if err != nil {
		return err
	}
	imdbID, err := valueobject.NewMovieID(cmd.IMDbID)
	if err != nil {
		return err
	}
	return app.CrawlAndSave(ctx, id, imdbID)
}
