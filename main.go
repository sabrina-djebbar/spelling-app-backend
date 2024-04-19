package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
	"github.com/sabrina-djebbar/spelling-app-backend/repo"
)

type CreateUserRequest struct {
	Username    string    `json:"username"`
	Password    string    `json:"password"`
	ParentCode  string    `json:"parent_code"`
	DateOfBirth time.Time `json:"date_of_birth"`
}
type User struct {
	Id          int
	Username    string
	ParentCode  string
	DateOfBirth time.Time
}

func main() {
	connStr := "postgres://postgres:secret@localhost:5432/spelling-app?sslmode=disable"

	db, err := sql.Open("postgres", connStr)

	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	userReq := repo.CreateUserParams{"user", time.Now(), "1234"}

	ctx := context.Background()
	queries := repo.New(*db)
	queries.CreateUser(ctx, userReq)
	/*
		// create tables
		createCredentialsTable(db)
		createUserTable(db)

		userReq := CreateUserRequest{"user", "password", "1234", time.Now()}
		user := createUser(db, userReq)
		fmt.Print("user ", user)*/
}

func createCredentialsTable(db *sql.DB) {
	_, err := db.Exec("CREATE EXTENSION IF NOT EXISTS pgcrypto")
	if err != nil {
		log.Fatal(err)
	}
	sqlQuery := `CREATE TABLE IF NOT EXISTS credentials (
		id SERIAL PRIMARY KEY,
		user_id SERIAL,
		password VARCHAR(50),
		FOREIGN KEY (user_id) REFERENCES users(id)
	);`
	_, err = db.Exec(sqlQuery)
	if err != nil {
		log.Fatal(err)
	}
}

func createUserTable(db *sql.DB) {
	sqlQuery := `CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		username VARCHAR(50) NOT NULL,
		parent_code VARCHAR(4),
		date_of_birth DATE,
		created timestamp DEFAULT NOW()
	);`

	_, err := db.Exec(sqlQuery)
	if err != nil {
		log.Fatal(err)
	}
}

func createUser(db *sql.DB, req CreateUserRequest) int {
	sqlUserQuery := `INSERT INTO users (username, date_of_birth, parent_code) VALUES ($1, $2, $3) RETURNING id, username`
	sqlCredentialQuery := `INSERT INTO credentials (user_id, password) VALUES($1,crypt($2,'crypt-des')));`
	var (
		id       int
		username string
	)
	err := db.QueryRow(sqlUserQuery, req.Username, req.DateOfBirth, req.ParentCode).Scan(&id, &username)
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec(sqlCredentialQuery, id, req.Password)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("user ", id, username)
	return id
}
