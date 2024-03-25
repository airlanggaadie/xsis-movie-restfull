package model

type Movie struct {
	// TODO: fill these fields
}

type ListMovieResponse struct {
	Data []Movie `json:"data"`
}

type MovieDetailResponse struct {
	Movie
}
