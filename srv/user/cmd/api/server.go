package api

import (
	"github.com/sabrina-djebbar/spelling-app-backend/lib/database"
	"github.com/sabrina-djebbar/spelling-app-backend/lib/shttp"
	"github.com/sabrina-djebbar/spelling-app-backend/lib/shttp/middleware"
	"github.com/sabrina-djebbar/spelling-app-backend/srv/user/internal/app"
	userRepo "github.com/sabrina-djebbar/spelling-app-backend/srv/user/internal/infrastructure"
	"github.com/sabrina-djebbar/spelling-app-backend/srv/user/internal/infrastructure/repo"
	"github.com/sabrina-djebbar/spelling-app-backend/srv/user/internal/rpc"
	"github.com/sabrina-djebbar/spelling-app-backend/srv/user/pkg/client"
	"github.com/spf13/cobra"
	"log"
)

var CMD = &cobra.Command{
	Use:   "api",
	Short: "User service implementation",
	Long:  "User service implements complete management of a user",
	RunE:  runE,
}
var logger log.Logger

func runE(cmd *cobra.Command, _ []string) error {
	db, err := database.New("user")
	if err != nil {
		logger.Fatal("unable to create postgres client", err)
	}

	queries := repo.New(db)
	repository := userRepo.NewRepo(*queries)
	var (
		a = app.New(repository)
		r = rpc.New(a)
	)

	router := shttp.New(cmd)
	router.RegisterMiddleware(middleware.NewLoggingMiddleware(cmd))
	router.RegisterHandler(client.GetUserPath, r.GetUser)
	router.RegisterHandler(client.CreateUserPath, r.CreateUser)
	router.RegisterHandler(client.ListUsersPath, r.ListUser)
	router.RegisterHandler(client.LoginPath, r.ListUser)

	return router.Listen(":8080")
}
