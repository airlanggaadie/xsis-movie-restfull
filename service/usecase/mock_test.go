package usecase

import (
	"context"
	"database/sql"
	"errors"
	"xsis/assignment-test/model"

	"github.com/google/uuid"
)

type movieRepositorySuccessMock struct{}

func (m movieRepositorySuccessMock) GetMoviesPaginate(ctx context.Context, offset, limit int) ([]model.Movie, int64, error) {
	return []model.Movie{}, 0, nil
}

func (m movieRepositorySuccessMock) GetMovie(ctx context.Context, id uuid.UUID) (model.Movie, error) {
	return model.Movie{}, nil
}

func (m movieRepositorySuccessMock) InsertMovie(ctx context.Context, movie model.Movie) (model.Movie, error) {
	idUUID, _ := uuid.Parse("fe9ba0d7-3f78-472d-9c75-5c5ade522b84")
	return model.Movie{
		Id:          idUUID,
		Title:       "test",
		Description: "description test",
		Rating:      5,
		Image:       "image test",
	}, nil
}

func (m movieRepositorySuccessMock) UpdateMovie(ctx context.Context, newMovie model.Movie) (model.Movie, error) {
	idUUID, _ := uuid.Parse("fe9ba0d7-3f78-472d-9c75-5c5ade522b84")
	return model.Movie{
		Id:          idUUID,
		Title:       "test",
		Description: "description test",
		Rating:      5,
		Image:       "image test",
	}, nil
}

func (m movieRepositorySuccessMock) DeleteMovie(ctx context.Context, id uuid.UUID) error {
	return nil
}

type movieRepositoryFailMock struct{}

func (m movieRepositoryFailMock) GetMoviesPaginate(ctx context.Context, offset, limit int) ([]model.Movie, int64, error) {
	return []model.Movie{}, 0, errors.New("failed to get movies")
}

func (m movieRepositoryFailMock) GetMovie(ctx context.Context, id uuid.UUID) (model.Movie, error) {
	return model.Movie{}, sql.ErrNoRows
}

func (m movieRepositoryFailMock) InsertMovie(ctx context.Context, movie model.Movie) (model.Movie, error) {
	return model.Movie{}, errors.New("failed to insert movie")
}

func (m movieRepositoryFailMock) UpdateMovie(ctx context.Context, newMovie model.Movie) (model.Movie, error) {
	return model.Movie{}, sql.ErrNoRows
}

func (m movieRepositoryFailMock) DeleteMovie(ctx context.Context, id uuid.UUID) error {
	return sql.ErrNoRows
}
