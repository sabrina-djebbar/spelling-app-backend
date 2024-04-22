package database

import (
	"database/sql"
	"strconv"

	_ "github.com/lib/pq"
)

// Database connection parameters
const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "secret"
)

func New(name string) (*sql.DB, error) {
	connStr := "postgres://" + user + ":" + password + "@" + host + ":" + strconv.Itoa(port) + "/" + name + "?sslmode=disable"
	//connStr := "host=" + host + " port=" + strconv.Itoa(port) + " user=" + user + " password=" + password + " dbname=" + name + " sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
