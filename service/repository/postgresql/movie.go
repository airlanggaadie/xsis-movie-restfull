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

func (m movieRepository) GetMovie(ctx context.Context, id uuid.UUID) (model.Movie, error) {
	var movie model.Movie
	if err := m.DB.QueryRowContext(ctx, queryGetMovieById, id).Scan(
		&movie.Id,
		&movie.Title,
		&movie.Description,
		&movie.Rating,
		&movie.Image,
		&movie.CreatedAt,
		&movie.UpdatedAt,
	); err != nil {
		return movie, fmt.Errorf("[postgresql][GetMovie] error query: %w", err)
	}

	return movie, nil
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

func (m movieRepository) UpdateMovie(ctx context.Context, newMovie model.Movie) (model.Movie, error) {
	tx, err := m.DB.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return model.Movie{}, fmt.Errorf("[postgresql][UpdateMovie] begin transaction error: %w", err)
	}
	defer tx.Rollback()

	var updatedMovie model.Movie
	if err := tx.QueryRowContext(ctx, queryUpdateMovie, newMovie.Id, newMovie.Title, newMovie.Description, newMovie.Rating, newMovie.Image).Scan(
		&updatedMovie.Id,
		&updatedMovie.Title,
		&updatedMovie.Description,
		&updatedMovie.Rating,
		&updatedMovie.Image,
		&updatedMovie.CreatedAt,
		&updatedMovie.UpdatedAt,
	); err != nil {
		return model.Movie{}, fmt.Errorf("[postgresql][UpdateMovie] execution error: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return model.Movie{}, fmt.Errorf("[postgresql][UpdateMovie] commit error: %w", err)
	}

	return updatedMovie, nil
}

func (m movieRepository) DeleteMovie(ctx context.Context, id uuid.UUID) error {
	tx, err := m.DB.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return fmt.Errorf("[postgresql][DeleteMovie] begin transaction error: %w", err)
	}
	defer tx.Rollback()

	_, err = tx.ExecContext(ctx, queryDeleteMovie, id)
	if err != nil {
		return fmt.Errorf("[postgresql][DeleteMovie] execution error: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("[postgresql][DeleteMovie] commit error: %w", err)
	}

	return nil
}
