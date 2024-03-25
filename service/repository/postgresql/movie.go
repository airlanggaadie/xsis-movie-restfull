package postgresql

import (
	"context"
	"database/sql"
	"fmt"
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

func (m movieRepository) GetMoviesPaginate(ctx context.Context, offset, limit int) ([]model.Movie, int64, error) {
	query, err := m.DB.QueryContext(ctx, queryGetMoviesPaginate, offset, limit)
	if err != nil {
		return nil, 0, fmt.Errorf("[postgresql][GetMoviesPaginate] error query: %v", err)
	}
	defer query.Close()

	var movies []model.Movie
	for query.Next() {
		var movie model.Movie
		if err := query.Scan(
			&movie.Id,
			&movie.Title,
			&movie.Description,
			&movie.Rating,
			&movie.Image,
			&movie.CreatedAt,
			&movie.UpdatedAt,
		); err != nil {
			return nil, 0, fmt.Errorf("[postgresql][GetMoviesPaginate] error scan: %v", err)
		}

		movies = append(movies, movie)
	}

	var total int64
	if err := m.DB.QueryRowContext(ctx, queryGetMoviesCount).Scan(&total); err != nil {
		return nil, 0, fmt.Errorf("[postgresql][GetMoviesPaginate] error count: %v", err)
	}

	return movies, total, nil
}

func (m movieRepository) GetMovie(id uuid.UUID) (model.Movie, error) {
	// TODO: do something
	return model.Movie{}, nil
}

func (m movieRepository) InsertMovie(ctx context.Context, movie model.Movie) (model.Movie, error) {
	tx, err := m.DB.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return model.Movie{}, fmt.Errorf("[postgresql][InsertMovie] begin transaction error: %w", err)
	}
	defer tx.Rollback()

	_, err = tx.ExecContext(ctx, queryInsertMovie, movie.Id, movie.Title, movie.Description, movie.Rating, movie.Image)
	if err != nil {
		return model.Movie{}, fmt.Errorf("[postgresql][InsertMovie] execution error: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return model.Movie{}, fmt.Errorf("[postgresql][InsertMovie] commit error: %w", err)
	}

	return movie, nil
}

func (m movieRepository) UpdateMovie(id uuid.UUID, movie model.Movie) (model.Movie, error) {
	// TODO: do something
	return model.Movie{}, nil
}

func (m movieRepository) DeleteMovie(id uuid.UUID) error {
	// TODO: do something
	return nil
}
