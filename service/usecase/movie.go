package usecase

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
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

func (m movie) GetMoviesPaginate(ctx context.Context, search string, page, limit int) (model.ListMovieResponse, error) {
	if page == 0 {
		page = 1
	}

	if limit == 0 {
		limit = 10
	}

	offset := (page - 1) * limit

	movies, total, err := m.movieRepository.GetMoviesPaginate(ctx, search, offset, limit)
	if err != nil {
		return model.ListMovieResponse{}, fmt.Errorf("[usecase][GetMoviePaginate] error get movies: %v", err)
	}

	return model.ListMovieResponse{
		Data:  movies,
		Total: total,
	}, nil
}

func (m movie) GetMovie(ctx context.Context, id uuid.UUID) (model.MovieDetailResponse, error) {
	movie, err := m.movieRepository.GetMovie(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.MovieDetailResponse{}, model.ErrNotFound
		}
		return model.MovieDetailResponse{}, fmt.Errorf("[usecase][GetMovie] error: %v", err)
	}

	return model.MovieDetailResponse{
		Movie: movie,
	}, nil
}

func (m movie) AddNewMovie(ctx context.Context, request model.NewMovieRequest) (model.MovieDetailResponse, error) {
	// prepare new movie
	newMovie, err := model.NewMovie(request)
	if err != nil {
		return model.MovieDetailResponse{}, fmt.Errorf("[usecase][AddNewMovie] error new movie: %v", err)
	}

	// add new movie to db
	newMovie, err = m.movieRepository.InsertMovie(ctx, newMovie)
	if err != nil {
		return model.MovieDetailResponse{}, fmt.Errorf("[usecase][AddNewMovie] error create movie: %v", err)
	}

	return model.MovieDetailResponse{
		Movie: newMovie,
	}, nil
}

func (m movie) UpdateMovie(ctx context.Context, id uuid.UUID, request model.NewMovieRequest) (model.MovieDetailResponse, error) {
	updatedMovie, err := m.movieRepository.UpdateMovie(ctx, model.NewUpdateMovie(id, request))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.MovieDetailResponse{}, model.ErrNotFound
		}

		return model.MovieDetailResponse{}, fmt.Errorf("[usecase][UpdateMovie] error update movie: %v", err)
	}

	return model.MovieDetailResponse{
		Movie: updatedMovie,
	}, nil
}

func (m movie) DeleteMovie(ctx context.Context, id uuid.UUID) error {
	movie, err := m.movieRepository.GetMovie(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.ErrNotFound
		}
		return fmt.Errorf("[usecase][DeleteMovie] error: %v", err)
	}

	err = m.movieRepository.DeleteMovie(ctx, movie.Id)
	if err != nil {
		return fmt.Errorf("[usecase][DeleteMovie] error delete: %v", err)
	}

	return nil
}
