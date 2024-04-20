package database

import (
	"database/sql"
	"fmt"
)

func New(name string) (*sql.DB, error) {
	connStr := fmt.Sprintf("postgres://postgres:secret@localhost:5432/%s?sslmode=disable", name)

	db, err := sql.Open("postgres", connStr)
	defer db.Close()
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
