package client

type CreateUserRequest struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	ParentCode  string `json:"parent_code"`
	DateOfBirth string `json:"date_of_birth"`
}

type CreateUserResponse struct {
	User *models.User `json:"user"`
}
