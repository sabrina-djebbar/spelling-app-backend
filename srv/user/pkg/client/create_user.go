package client

import (
	"github.com/sabrina-djebbar/spelling-app-backend/srv/user/pkg/models"
	"time"
)

type CreateUserRequest struct {
	Username    string    `json:"username"`
	Password    string    `json:"password"`
	ParentCode  string    `json:"parent_code"`
	DateOfBirth time.Time `json:"date_of_birth"`
}

type CreateUserResponse struct {
	User *models.User `json:"user"`
}
