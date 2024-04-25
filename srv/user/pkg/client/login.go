package client

import "github.com/sabrina-djebbar/spelling-app-backend/srv/user/pkg/models"

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	User *models.User `json:"user"`
}
