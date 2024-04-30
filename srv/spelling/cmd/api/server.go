package api

import (
	"github.com/sabrina-djebbar/spelling-app-backend/lib/database"
	"github.com/sabrina-djebbar/spelling-app-backend/lib/shttp"
	"github.com/sabrina-djebbar/spelling-app-backend/srv/spelling/internal/app"
	spellingRepo "github.com/sabrina-djebbar/spelling-app-backend/srv/spelling/internal/infrastructure"
	"github.com/sabrina-djebbar/spelling-app-backend/srv/spelling/internal/infrastructure/repo"
	"github.com/sabrina-djebbar/spelling-app-backend/srv/spelling/internal/rpc"
	"github.com/sabrina-djebbar/spelling-app-backend/srv/spelling/pkg/client"
	"github.com/spf13/cobra"
	"log"
)

var CMD = &cobra.Command{
	Use:   "api",
	Short: "Spelling service implementation",
	Long:  "Spelling service implements complete management of a user",
	RunE:  runE,
}
var logger log.Logger

func runE(cmd *cobra.Command, _ []string) error {
	db, err := database.New(&logger, "spelling")
	if err != nil {
		logger.Fatal("unable to create postgres client", err)
	}

	queries := repo.New(db)
	repository := spellingRepo.NewRepo(*queries)
	var (
		a = app.New(repository)
		r = rpc.New(a)
	)

	router := shttp.New(cmd)
	router.RegisterHandler(client.CreateSpellingWordPath, r.CreateSpellingWord)
	router.RegisterHandler(client.CreateSpellingSetPath, r.CreateSpellingSet)
	router.RegisterHandler(client.ListSpellingSetsPath, r.ListSpellingSets)
	router.RegisterHandler(client.CreateSpellingAttemptPath, r.CreateSpellingAttempt)
	router.RegisterHandler(client.ListSpellingExerciseByUserPath, r.ListSpellingExerciseByUser)
	return router.Listen("8081")
}
