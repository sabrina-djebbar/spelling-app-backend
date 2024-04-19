package main

import (
	"context"
	"crypto/rand"
	"database/sql"
	"encoding/base64"
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

func Generate(resourceType string) string {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		panic(fmt.Sprintf("firestore: crypto/rand.Read error: %v", err))
	}

	r := base64.RawURLEncoding.EncodeToString(b)
	return fmt.Sprintf("%s_%s", r, resourceType)
}

// https://www.youtube.com/watch?v=x_N2VjGQKr4
func main() {
	ctx := context.Background()

	connStr := "postgres://postgres:secret@localhost:5432/spelling-app?sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	var dob sql.NullTime
	userId := Generate("user")
	userReq := repo.CreateUserParams{ID: userId, Username: "user3", DateOfBirth: dob, ParentCode: "1234"}

	queries := repo.New(db)
	_, err = queries.CreateUser(ctx, userReq)
	if err != nil {
		log.Fatal(err)
	}
	err = queries.CreateCredentials(ctx, repo.CreateCredentialsParams{ID: Generate("credential"), UserID: userId, Crypt: "password"})
	if err != nil {
		log.Fatal(err)
	}
	/*
		// create tables
		createCredentialsTable(db)
		createUserTable(db)

		userReq := CreateUserRequest{"user", "password", "1234", time.Now()}
		user := createUser(db, userReq)
		fmt.Print("user ", user)*/
}
