package postgresql

const (
	queryGetMoviesPaginate = `SELECT id, title, description, rating, image, created_at, updated_at FROM movie OFFSET $1 LIMIT $2;`
	queryGetMoviesCount    = `SELECT count(*) FROM movie;`
	queryInsertMovie       = `INSERT INTO movie (id, title, description, rating, image) VALUES ($1, $2, $3, $4, $5);`
)
