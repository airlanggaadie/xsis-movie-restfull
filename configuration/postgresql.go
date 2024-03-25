package configuration

import (
	"context"
	"database/sql"
	"log"
	"os"
	"time"

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

func (c *configuration) migrate() *configuration {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	tx, err := c.DB.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable, ReadOnly: false})
	if err != nil {
		log.Fatalf("[configuration][migrate] begin transaction error: %v", err)
	}

	queries := []string{
		`CREATE TABLE IF NOT EXISTS "movie" (
			id                  UUID PRIMARY KEY NOT NULL,
			title               VARCHAR(50) NOT NULL,
			description			TEXT DEFAULT '',
			rating 				DECIMAL DEFAULT 0,
			image				TEXT DEFAULT '',
			created_at          TIMESTAMP NOT NULL,
			updated_at          TIMESTAMP NOT NULL
		);`,
	}

	for i, query := range queries {
		_, err = tx.ExecContext(
			ctx,
			query,
		)
		if err != nil {
			tx.Rollback()
			log.Fatalf("[configuration][migrate] execution [%d] error: %v", i, err)
		}

	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		log.Fatalf("[configuration][migrate] commit error: %v", err)
	}

	return c
}
