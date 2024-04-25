package app

import "github.com/sabrina-djebbar/spelling-app-backend/srv/spelling/internal/infrastructure/repo"

type App interface {
	GetUser(ctx context.Context, userID string) (*models.User, error)
	CreateUser(ctx context.Context, req CreateUserRequest) (*models.User, error)
	EditUser(ctx context.Context, req EditUserRequest) (*models.User, error)
	EditParentDetails(ctx context.Context, req EditParentDetailsRequest) (*models.User, error)
	Login(ctx context.Context, req LoginRequest) (*models.User, error)
	Logout(ctx context.Context, req LogoutRequest) error
}

type app struct {
	repository repo.UserRepository
}

func New(
	repository repo.UserRepository,
) App {
	return &app{
		repository,
	}
}
