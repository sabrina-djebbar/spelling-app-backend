package user

import (
	"context"
	"fmt"
	"github.com/sabrina-djebbar/spelling-app-backend/lib/database"
	"github.com/sabrina-djebbar/spelling-app-backend/srv/user/internal/app"
	"github.com/sabrina-djebbar/spelling-app-backend/srv/user/internal/rpc"
	"github.com/sabrina-djebbar/spelling-app-backend/srv/user/pkg/client"

	userRepo "github.com/sabrina-djebbar/spelling-app-backend/srv/user/internal/infrastructure"
	"github.com/sabrina-djebbar/spelling-app-backend/srv/user/internal/infrastructure/repo"
	"log"
)

func main(logger *log.Logger) {
	db, err := database.New(logger, "user")
	if err != nil {
		logger.Fatal("unable to create postgres client", err)
	}
	defer db.Close()

	ctx := context.Background()
	queries := repo.New(db)
	repository := userRepo.NewRepo(*queries)
	var (
		a = app.New(repository)
		r = rpc.New(a)
	)

	var params = client.EditUserRequest{
		UserID:      "VVllNSUIGg03MrmrDWW52fJszDi9ITS2Ly6uWp5Okdc_user",
		DateOfBirth: "2002-01-02T00:00:00Z",
	}
  
	// router := shttp.New(cmd)
	//router.RegisterMiddleware(middleware.NewLoggingMiddleware(cmd))
	//router.RegisterHandler(client.GetUserPath, r.GetUser)
	//router.RegisterHandler(client.CreateUserPath, r.CreateUser)
	//	router.RegisterHandler(client.ListUsersPath, r.ListUser)
	//	router.RegisterHandler(client.LoginPath, r.ListUser)

	// return router.Listen(":8080")

	res, err := r.EditUser(ctx, params)

	if err != nil {
		fmt.Println("unable to login user", err)
	}
	fmt.Println("response", res)
}
