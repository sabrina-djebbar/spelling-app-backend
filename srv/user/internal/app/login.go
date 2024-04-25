package app

import (
	"context"
	userRepo "github.com/sabrina-djebbar/spelling-app-backend/srv/user/internal/infrastructure"
	"github.com/sabrina-djebbar/spelling-app-backend/srv/user/pkg/client"
	"github.com/sabrina-djebbar/spelling-app-backend/srv/user/pkg/models"
)

func (a *app) Login(ctx context.Context, req client.LoginRequest) (*models.User, error) {
	user, err := a.repository.FindByUsername(ctx, req.Username)
	if err != nil {
		return nil, err
	}

	_, err = a.repository.FindCredentials(ctx, userRepo.FindCredentialParams{UserID: user.ID, Crypt: req.Password})
	if err != nil {
		return nil, err
	}
	return &models.User{
		ID:         user.ID,
		Username:   user.Username,
		Birthday:   user.DateOfBirth.Time,
		ParentCode: user.ParentCode,
	}, nil
}
