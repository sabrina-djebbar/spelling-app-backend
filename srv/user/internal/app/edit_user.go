package app

import (
	"context"
	"github.com/sabrina-djebbar/spelling-app-backend/lib/database"
	"github.com/sabrina-djebbar/spelling-app-backend/lib/serr"
	"github.com/sabrina-djebbar/spelling-app-backend/srv/user/pkg/client"
	"github.com/sabrina-djebbar/spelling-app-backend/srv/user/pkg/models"
	"time"
)

func (a *app) EditUser(ctx context.Context, req client.EditUserRequest) (*models.User, error) {
	user, err := a.GetUser(ctx, req.UserID)
	if err != nil {
		return nil, err
	}

	if req.Username != "" {
		user.Username = req.Username
	}
	if req.DateOfBirth != "" {
		dob, err := time.Parse(time.DateOnly, req.DateOfBirth)
		if err != nil {
			return nil, serr.Wrap(err, serr.WithMessage("Error parsing time"))
		}
		user.DateOfBirth = dob
	}

	u, err := a.repository.EditUser(ctx, *user)
	if err != nil {
		return nil, err
	}

	return &models.User{
		ID:          u.ID,
		Username:    u.Username,
		DateOfBirth: database.SQLNullTimeToTime(u.DateOfBirth),
		ParentCode:  u.ParentCode,
	}, nil
}
