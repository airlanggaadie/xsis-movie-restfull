package service

import (
	"context"
	"xsis/assignment-test/model"

	"github.com/google/uuid"
)

type MovieUsecase interface {
	GetMoviesPaginate() (model.ListMovieResponse, error)
	GetMovie(id uuid.UUID) (model.MovieDetailResponse, error)

	// AddNewMovie adds a new movie to the list of movies
	AddNewMovie(ctx context.Context, movie model.AddNewMovieRequest) (model.MovieDetailResponse, error)
	UpdateMovie(id uuid.UUID, movie model.Movie) (model.MovieDetailResponse, error)
	DeleteMovie(id uuid.UUID) error
}
