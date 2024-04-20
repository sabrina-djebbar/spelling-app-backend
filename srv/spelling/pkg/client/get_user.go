package client

type GetUserRequest struct {
	UserID string `json:"user_id"`
}

type GetUserResponse struct {
	User *models.User `json:"user"`
}
