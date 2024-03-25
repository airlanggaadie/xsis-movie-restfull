package configuration

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func (c *configuration) initPostgreSql() *configuration {
	dbURL, ok := os.LookupEnv("DATABASE_URL")
	if !ok {
		dbURL = "postgres://postgres:postgres@localhost:5432/xsis?sslmode=disable"
	}

	// setup the database connection
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("[configuration][initPostgreSql] failed to connect to database: %v", err)
	}

	c.DB = db

	return c
}
