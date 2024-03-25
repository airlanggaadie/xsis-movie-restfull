package service

import (
	"xsis/assignment-test/model"

	"github.com/google/uuid"
)

type MovieRepository interface {
	GetMoviesPaginate() ([]model.Movie, error)
	GetMovie(id uuid.UUID) (model.Movie, error)
	CreateMovie(movie model.Movie) (model.Movie, error)
	UpdateMovie(id uuid.UUID, movie model.Movie) (model.Movie, error)
	DeleteMovie(id uuid.UUID) error
}
