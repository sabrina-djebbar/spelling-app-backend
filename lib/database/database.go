package database

import (
	"database/sql"
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

var logger = log.Logger{}

func New(name string) (*sql.DB, error) {
	connStr := "postgres://" + user + ":" + password + "@" + host + ":" + strconv.Itoa(port) + "/" + name + "?sslmode=disable"
	//connStr := "host=" + host + " port=" + strconv.Itoa(port) + " user=" + user + " password=" + password + " dbname=" + name + " sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		logger.Println("error initialising db " + name + "\n " + err.Error() + "\n")
		return nil, err
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			logger.Println("error closing db " + name + "\n " + err.Error() + "\n")
		}
	}(db)

	if err = db.Ping(); err != nil {
		logger.Println("error pinging db " + name + "\n " + err.Error() + "\n")
		return nil, err
	}
	logger.Println("connected to db " + name + "\n")
	return db, nil
}
