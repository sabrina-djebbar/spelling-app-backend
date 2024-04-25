package client

import "github.com/sabrina-djebbar/spelling-app-backend/srv/user/pkg/models"

type ListUsersRequest struct {
}

type ListUsersResponse struct {
	Users []models.User `json:"users"`
}
