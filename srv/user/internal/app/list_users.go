package app

import (
	"context"
	"github.com/sabrina-djebbar/spelling-app-backend/srv/user/pkg/models"
)

func (a *app) ListUsers(ctx context.Context) ([]models.User, error) {
	u, err := a.repository.ListUsers(ctx)
	if err != nil {
		return nil, err
	}

	var users []models.User

	for _, user := range *u {
		users = append(users, models.User{
			ID:          user.ID,
			Username:    user.Username,
			DateOfBirth: user.DateOfBirth.Time, // Access Time from sql.NullTime
			ParentCode:  user.ParentCode,
		})
	}

	return users, nil
}
