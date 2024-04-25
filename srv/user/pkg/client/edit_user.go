package client

import "github.com/sabrina-djebbar/spelling-app-backend/srv/user/pkg/models"

type EditUserRequest struct {
	UserID      string `json:"user_id"`
	Username    string `json:"username"`
	DateOfBirth string `json:"date_of_birth"`
}

type EditUserResponse struct {
	User *models.User `json:"user"`
}
