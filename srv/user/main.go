package user

import (
	"context"
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/sabrina-djebbar/spelling-app-backend/srv/user/internal/repo"
)

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
}
