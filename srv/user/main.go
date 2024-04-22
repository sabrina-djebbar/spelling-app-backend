package user

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/sabrina-djebbar/spelling-app-backend/lib/database"

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
	/* var (
		a = app.New(repository)
		r = rpc.New(a)
	)*/
	var params = userRepo.CreateUserParams{
		Username:    "Username",
		DateOfBirth: time.Time{},
		ParentCode:  "1234",
		Password:    "password",
	}
	// router := shttp.New()
	//router.RegisterMiddleware(middleware.NewLoggingMiddleware(cmd))
	//router.RegisterHandler(client.GetUserPath, r.GetUser)
	//router.RegisterHandler(client.CreateUserPath, r.CreateUser)
	//	router.RegisterHandler(client.ListUsersPath, r.ListUser)
	//	router.RegisterHandler(client.LoginPath, r.ListUser)

	// return router.Listen(":8080")
	_, err = repository.CreateUser(ctx, params)
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
