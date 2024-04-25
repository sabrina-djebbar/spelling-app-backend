package app

import (
	"context"
	repo "github.com/sabrina-djebbar/spelling-app-backend/srv/user/internal/infrastructure"
	"github.com/sabrina-djebbar/spelling-app-backend/srv/user/pkg/client"
	"github.com/sabrina-djebbar/spelling-app-backend/srv/user/pkg/models"
)

type App interface {
	GetUser(ctx context.Context, userID string) (*models.User, error)
	CreateUser(ctx context.Context, req client.CreateUserRequest) (*models.User, error)
	EditUser(ctx context.Context, req client.EditUserRequest) (*models.User, error)
	EditParentDetails(ctx context.Context, req client.EditParentCodeRequest) (*models.User, error)
	Login(ctx context.Context, req client.LoginRequest) (*models.User, error)
	ListUsers(ctx context.Context) ([]models.User, error)
}

type app struct {
	repository repo.Repository
}

func New(
	repository repo.Repository,
) App {
	return &app{
		repository,
	}
}
