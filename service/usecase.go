package service

import (
	"xsis/assignment-test/model"

	"github.com/google/uuid"
)

type MovieUsecase interface {
	GetMoviesPaginate() (model.ListMovieResponse, error)
	GetMovie(id uuid.UUID) (model.MovieDetailResponse, error)
	CreateMovie(movie model.Movie) (model.MovieDetailResponse, error)
	UpdateMovie(id uuid.UUID, movie model.Movie) (model.MovieDetailResponse, error)
	DeleteMovie(id uuid.UUID) error
}
