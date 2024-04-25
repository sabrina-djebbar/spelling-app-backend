package api

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/sabrina-djebbar/spelling-app-backend/lib/database"
	"github.com/sabrina-djebbar/spelling-app-backend/srv/user/client"
	userRepo "github.com/sabrina-djebbar/spelling-app-backend/srv/user/internal/infrastructure"
	"github.com/sabrina-djebbar/spelling-app-backend/srv/user/internal/infrastructure/repo"
)

var CMD = &cobra.Command{
	Use:   "api",
	Short: "User service implementation",
	Long:  "User service implements complete management of a user",
	RunE:  runE,
}

func RunE() {
	ctx := context.Background()
	db, err := database.New(ctx, "user")
	queries := repo.New(db)

	repository := userRepo.NewRepo(queries)
	var (
		a = app.New(repository)
		r = rpc.New(a)
	)
	srv := gin.Default()
	srv.GET(client.getUserPath, r.GetUser)
	srv.POST(client.createUserPath, r.CreateUser)
	srv.POST(client.loginPath, r.Login)
	srv.POST(client.logoutPath, r.Logout)
	srv.PUT(client.editUserPath, r.EditUser)
	srv.PUT(client.editParentDetailsPath, r.EditParentDetails)

	srv.Run(":8080")
}
