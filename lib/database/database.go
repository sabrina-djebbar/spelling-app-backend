package database

import (
	"context"
	"database/sql"
	"fmt"
)

func New(ctx context.Context, name string) (*sql.DB, error) {
	connStr := fmt.Sprintf("postgres://postgres:secret@localhost:5432/%s?sslmode=disable", name)

	db, err := sql.Open("postgres", connStr)
	defer db.Close()
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return err, nil
	}
	return db, nil
}
