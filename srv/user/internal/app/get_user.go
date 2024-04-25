package app

import (
	"context"
	"github.com/sabrina-djebbar/spelling-app-backend/srv/user/pkg/models"
)

func (a *app) GetUser(ctx context.Context, userID string) (*models.User, error) {
	u, err := a.repository.GetUser(ctx, userID)
	if err != nil {
		return nil, err
	}

	return &models.User{ID: u.ID, Username: u.Username, DateOfBirth: u.DateOfBirth.Time, ParentCode: u.ParentCode}, nil
}
