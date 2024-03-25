package service

import (
	"context"
	"xsis/assignment-test/model"

	"github.com/google/uuid"
)

type MovieUsecase interface {
	// GetMoviesPaginate returns a list of movies with total data
	GetMoviesPaginate(ctx context.Context, page, limit int) (model.ListMovieResponse, error)

	// GetMovie returns a single movie
	GetMovie(ctx context.Context, id uuid.UUID) (model.MovieDetailResponse, error)

	// AddNewMovie adds a new movie to the list of movies
	AddNewMovie(ctx context.Context, movie model.AddNewMovieRequest) (model.MovieDetailResponse, error)
	UpdateMovie(id uuid.UUID, movie model.Movie) (model.MovieDetailResponse, error)
	DeleteMovie(id uuid.UUID) error
}
