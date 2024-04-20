package app

import (
	"context"
	"github.com/sabrina-djebbar/spelling-app-backend/lib/id"
	userRepo "github.com/sabrina-djebbar/spelling-app-backend/srv/user/internal/infrastructure"
	"github.com/sabrina-djebbar/spelling-app-backend/srv/user/internal/infrastructure/repo"
	"github.com/sabrina-djebbar/spelling-app-backend/srv/user/pkg/client"
	"time"
)

type CreateUserParams struct {
	ID          string
	Username    string
	DateOfBirth time.Time
	ParentCode  string
	Password    string
}

func (a *app) CreateUser(ctx context.Context, req client.CreateUserRequest) (*repo.User, error) {

	userId := id.Generate("user")
	user := &CreateUserParams{
		ID:          userId,
		Username:    req.Username,
		DateOfBirth: req.DateOfBirth,
		ParentCode:  req.ParentCode,
		Password:    req.Password,
	}

	u, err := a.repository.CreateUser(ctx, userRepo.CreateUserParams(*user))
	if err != nil {
		return nil, err
	}

	return u, nil
}
