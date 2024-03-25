package usecase

import (
	"context"
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

func (m movie) GetMoviesPaginate(ctx context.Context, page, limit int) (model.ListMovieResponse, error) {
	if page == 0 {
		page = 1
	}

	if limit == 0 {
		limit = 10
	}

	offset := (page - 1) * limit

	movies, total, err := m.movieRepository.GetMoviesPaginate(ctx, offset, limit)
	if err != nil {
		return model.ListMovieResponse{}, fmt.Errorf("[usecase][GetMoviePaginate] error get movies: %v", err)
	}

	return model.ListMovieResponse{
		Data:  movies,
		Total: total,
	}, nil
}

func (m movie) GetMovie(id uuid.UUID) (model.MovieDetailResponse, error) {
	// TODO: do something
	return model.MovieDetailResponse{}, nil
}

func (m movie) AddNewMovie(ctx context.Context, request model.AddNewMovieRequest) (model.MovieDetailResponse, error) {
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

func (m movie) UpdateMovie(id uuid.UUID, movie model.Movie) (model.MovieDetailResponse, error) {
	// TODO: do something
	return model.MovieDetailResponse{}, nil
}

func (m movie) DeleteMovie(id uuid.UUID) error {
	// TODO: do something
	return nil
}
