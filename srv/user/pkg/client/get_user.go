package client

import "github.com/sabrina-djebbar/spelling-app-backend/srv/user/pkg/models"

type GetUserRequest struct {
	UserID string `json:"user_id"`
}

type GetUserResponse struct {
	User *models.User `json:"user"`
}
