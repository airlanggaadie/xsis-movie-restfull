package service

import (
	"context"
	"xsis/assignment-test/model"

	"github.com/google/uuid"
)

type MovieRepository interface {
	// GetMoviesPaginate return a paginated list of movies
	GetMoviesPaginate(ctx context.Context, search string, offset, limit int) ([]model.Movie, int64, error)

	// GetMovie returns a single movie
	GetMovie(ctx context.Context, id uuid.UUID) (model.Movie, error)

	// InsertMovie inserts a new movie in the database
	InsertMovie(ctx context.Context, movie model.Movie) (model.Movie, error)

	// UpdateMovie updates an existing movie in the database
	UpdateMovie(ctx context.Context, newMovie model.Movie) (model.Movie, error)

	// DeleteMovie deletes a movie from the database
	DeleteMovie(ctx context.Context, id uuid.UUID) error
}
