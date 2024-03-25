package model

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Movie struct {
	Id          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Rating      float64   `json:"rating"`
	Image       string    `json:"image"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// NewMovie creates a new Movie with the given AddNewMovieRequest
func NewMovie(request AddNewMovieRequest) (Movie, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return Movie{}, err
	}

	return Movie{
		Id:          id,
		Title:       request.Title,
		Description: request.Description,
		Rating:      request.Rating,
		Image:       request.Image,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}, nil
}

type ListMovieResponse struct {
	Data  []Movie `json:"data"`
	Total int64   `json:"total"`
}

type MovieDetailResponse struct {
	Movie
}

type AddNewMovieRequest struct {
	Title       string  `json:"title"`
	Description string  `json:"description,omitempty"`
	Rating      float64 `json:"rating,omitempty"`
	Image       string  `json:"image,omitempty"`
}

func (r *AddNewMovieRequest) Validate() error {
	if r.Title == "" {
		return fmt.Errorf("title cannot be empty")
	}

	if r.Rating < 0 || r.Rating > 10 {
		return fmt.Errorf("rating must be between 0 and 10")
	}

	return nil
}
