package postgresql

import (
	"database/sql"
	"xsis/assignment-test/model"
	"xsis/assignment-test/service"

	"github.com/google/uuid"
)

type movieRepository struct {
	DB *sql.DB
}

func NewMovieRepository(db *sql.DB) service.MovieRepository {
	return movieRepository{
		DB: db,
	}
}

func (m movieRepository) GetMoviesPaginate() ([]model.Movie, error) {
	// TODO: do something
	return []model.Movie{}, nil
}

func (m movieRepository) GetMovie(id uuid.UUID) (model.Movie, error) {
	// TODO: do something
	return model.Movie{}, nil
}

func (m movieRepository) CreateMovie(movie model.Movie) (model.Movie, error) {
	// TODO: do something
	return model.Movie{}, nil
}

func (m movieRepository) UpdateMovie(id uuid.UUID, movie model.Movie) (model.Movie, error) {
	// TODO: do something
	return model.Movie{}, nil
}

func (m movieRepository) DeleteMovie(id uuid.UUID) error {
	// TODO: do something
	return nil
}
