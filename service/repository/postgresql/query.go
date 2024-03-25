package postgresql

const (
	queryGetMoviesPaginate = `SELECT id, title, description, rating, image, created_at, updated_at FROM movie OFFSET $1 LIMIT $2;`
	queryGetMoviesCount    = `SELECT count(*) FROM movie;`
	queryGetMovieById      = `SELECT id, title, description, rating, image, created_at, updated_at FROM movie WHERE id = $1;`
	queryInsertMovie       = `INSERT INTO movie (id, title, description, rating, image) VALUES ($1, $2, $3, $4, $5);`
	queryUpdateMovie       = `UPDATE movie SET title = $2, description = $3, rating = $4, image = $5, updated_at = NOW() WHERE id = $1 RETURNING id, title, description, rating, image, created_at, updated_at;`
	queryDeleteMovie       = `DELETE FROM movie WHERE id = $1;`
)
