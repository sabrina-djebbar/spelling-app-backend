package app

import (
	"context"
	userRepo "github.com/sabrina-djebbar/spelling-app-backend/srv/user/internal/infrastructure"
	"github.com/sabrina-djebbar/spelling-app-backend/srv/user/pkg/client"
	"github.com/sabrina-djebbar/spelling-app-backend/srv/user/pkg/models"
	"time"
)

type CreateUserParams struct {
	ID          string
	Username    string
	DateOfBirth time.Time
	ParentCode  string
	Password    string
}

func (a *app) CreateUser(ctx context.Context, req client.CreateUserRequest) (*models.User, error) {

	// userId := id.Generate("user")
	user := userRepo.CreateUserParams{

		Username:    req.Username,
		DateOfBirth: req.DateOfBirth,
		ParentCode:  req.ParentCode,
		Password:    req.Password,
	}

	u, err := a.repository.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return &models.User{
		ID:         u.ID,
		Username:   u.Username,
		Birthday:   u.DateOfBirth.Time,
		ParentCode: u.ParentCode,
	}, nil
}
