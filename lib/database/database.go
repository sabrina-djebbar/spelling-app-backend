package database

import (
	"database/sql"
	"fmt"
	"github.com/sabrina-djebbar/spelling-app-backend/lib/serr"
	"log"
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

func New(logger *log.Logger, name string) (*sql.DB, error) {
	connStr := "postgres://" + user + ":" + password + "@" + host + ":" + strconv.Itoa(port) + "/" + name + "?sslmode=disable"
	// connStr := "host=" + host + " port=" + strconv.Itoa(port) + " user=" + user + " password=" + password + " dbname=" + name + " sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		logger.Println("error initialising db " + name + "\n " + err.Error() + "\n")
		return nil, fmt.Errorf("failed to initialise db: %w", err)
	}

	if err = db.Ping(); err != nil {
		logger.Println("error pinging db " + name + "\n " + err.Error() + "\n")
		return nil, serr.Wrap(err, serr.WithMessage("unable to open database"))
	}

	return db, nil
}
