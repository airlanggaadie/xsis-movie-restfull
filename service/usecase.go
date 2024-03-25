package service

import (
	"context"
	"xsis/assignment-test/model"

	"github.com/google/uuid"
)

type MovieUsecase interface {
	// GetMoviesPaginate returns a list of movies with total data
	GetMoviesPaginate(ctx context.Context, page, limit int) (model.ListMovieResponse, error)

	// GetMovie returns a single movie. it will return [ErrNotFound] when the movie is not exist in the database
	GetMovie(ctx context.Context, id uuid.UUID) (model.MovieDetailResponse, error)

	// AddNewMovie adds a new movie to the list of movies
	AddNewMovie(ctx context.Context, request model.NewMovieRequest) (model.MovieDetailResponse, error)

	// UpdateMovie updates an existing movie in the database. it will return [ErrNotFound] when the movie is not exist in the database
	UpdateMovie(ctx context.Context, id uuid.UUID, request model.NewMovieRequest) (model.MovieDetailResponse, error)

	// DeleteMovie deletes a movie from the database. it will return [ErrNotFound] when the movie is not exist in the database
	DeleteMovie(ctx context.Context, id uuid.UUID) error
}
