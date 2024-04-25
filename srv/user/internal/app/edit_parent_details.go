package app

import (
	"context"
	"github.com/sabrina-djebbar/spelling-app-backend/srv/user/pkg/client"
	"github.com/sabrina-djebbar/spelling-app-backend/srv/user/pkg/models"
)

func (a *app) EditParentDetails(ctx context.Context, req client.EditParentCodeRequest) (*models.User, error) {
	u, err := a.repository.EditParentCode(ctx, req.UserID, req.ParentCode)
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
