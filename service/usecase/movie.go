package usecase

import (
	"xsis/assignment-test/model"
	"xsis/assignment-test/service"

	"github.com/google/uuid"
)

type movie struct {
	movieRepository service.MovieRepository
}

func NewMovieUsecase(movieRepository service.MovieRepository) service.MovieUsecase {
	return movie{
		movieRepository: movieRepository,
	}
}

func (m movie) GetMoviesPaginate() (model.ListMovieResponse, error) {
	// TODO: do something
	return model.ListMovieResponse{}, nil
}

func (m movie) GetMovie(id uuid.UUID) (model.MovieDetailResponse, error) {
	// TODO: do something
	return model.MovieDetailResponse{}, nil
}

func (m movie) CreateMovie(movie model.Movie) (model.MovieDetailResponse, error) {
	// TODO: do something
	return model.MovieDetailResponse{}, nil
}

func (m movie) UpdateMovie(id uuid.UUID, movie model.Movie) (model.MovieDetailResponse, error) {
	// TODO: do something
	return model.MovieDetailResponse{}, nil
}

func (m movie) DeleteMovie(id uuid.UUID) error {
	// TODO: do something
	return nil
}
