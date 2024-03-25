package postgresql

const (
	queryInsertMovie = `INSERT INTO movie (id, title, description, rating, image) VALUES ($1, $2, $3, $4, $5);`
)
