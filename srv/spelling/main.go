package main

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/sabrina-djebbar/spelling-app-backend/srv/spelling/internal/app"
	spellingRepo "github.com/sabrina-djebbar/spelling-app-backend/srv/spelling/internal/infrastructure"
	"github.com/sabrina-djebbar/spelling-app-backend/srv/spelling/internal/infrastructure/repo"
	"github.com/sabrina-djebbar/spelling-app-backend/srv/spelling/internal/rpc"
	"github.com/sabrina-djebbar/spelling-app-backend/srv/spelling/pkg/client"
	"log"
)

func main() {
	ctx := context.Background()

	connStr := "postgres://postgres:secret@localhost:5432/spelling?sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	queries := repo.New(db)
	repository := spellingRepo.NewRepo(*queries)
	var (
		a = app.New(repository)
		r = rpc.New(a)
		// words = []string{"aQke1e5VqmrEaFVi14N62WEFkN26jVAZ8gxpr9hq8pY_word"}
	)
	req := client.ListSpellingExercisesByUserRequest{
		UserID: "VVllNSUIGg03MrmrDWW52fJszDi9ITS2Ly6uWp5Okdc_user",
	}
	res, err := r.ListSpellingExerciseByUser(ctx, req)
	if err != nil {
		log.Fatal(err)
	}
	for _, e := range res.Exercises {
		fmt.Println(e)
	}
}
