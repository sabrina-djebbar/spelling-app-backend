package client

import "github.com/sabrina-djebbar/spelling-app-backend/srv/user/pkg/models"

type EditParentCodeRequest struct {
	UserID     string `json:"user_id""`
	ParentCode string `json:"parent_code"`
}

type EditParentCodeResponse struct {
	User models.User `json:"user"`
}
