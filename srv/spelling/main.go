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
	/*	req := client.CreateSpellingAttemptRequest{
			AttemptID:     "Cqa78R_bXCo8UJAke7rTFaGoOn8dJ3JrZ7VrjVJkxWQ_exercise",
			UserID:        "VVllNSUIGg03MrmrDWW52fJszDi9ITS2Ly6uWp5Okdc_user",
			SetID:         "8UbqpeH79u2rrMO8T7s9OVvAZVxj3glgmzVT4I7hv5w_set",
			WordID:        "COaDsZepgj7TzPB0BpNXmClBIida5ioY5XNxR2XhJbw_word",
			Spelling:      "father",
			Score:         7,
			NumOfAttempts: 1,
			LastAttempt:   time.Date(2024, 4, 30, 12, 0, 0, 0, time.UTC),
		}
	*/
	res, err := r.ListSpellingSets(ctx, client.ListSpellingSetsRequest{})
	if err != nil {
		log.Fatal(err)
	}
	for _, i2 := range res.Sets {
		fmt.Println(i2)
	}

}
