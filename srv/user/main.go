package user

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/sabrina-djebbar/spelling-app-backend/lib/database"
	"github.com/sabrina-djebbar/spelling-app-backend/srv/user/internal/app"
	"github.com/sabrina-djebbar/spelling-app-backend/srv/user/internal/rpc"
	"github.com/sabrina-djebbar/spelling-app-backend/srv/user/pkg/client"

	userRepo "github.com/sabrina-djebbar/spelling-app-backend/srv/user/internal/infrastructure"
	"github.com/sabrina-djebbar/spelling-app-backend/srv/user/internal/infrastructure/repo"
	"log"
	"time"
)

func Main(logger *log.Logger) {
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
	var params = client.CreateUserRequest{
		Username:    "user_2",
		DateOfBirth: time.Date(2002, time.July, 13, 00, 00, 00, 0, time.UTC),
		ParentCode:  "1234",
		Password:    "password",
	}
	// router := shttp.New(cmd)
	//router.RegisterMiddleware(middleware.NewLoggingMiddleware(cmd))
	//router.RegisterHandler(client.GetUserPath, r.GetUser)
	//router.RegisterHandler(client.CreateUserPath, r.CreateUser)
	//	router.RegisterHandler(client.ListUsersPath, r.ListUser)
	//	router.RegisterHandler(client.LoginPath, r.ListUser)

	// return router.Listen(":8080")
	_, err = r.CreateUser(ctx, params)
	if err != nil {
		fmt.Println("unable to create user", err)
	}
	productList, err := queries.ListUsers(ctx)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			fmt.Println("no entries found, " + err.Error())
		}
		fmt.Println(err)
	}
	for _, element := range productList {
		fmt.Println(element.Username)
	}
}
