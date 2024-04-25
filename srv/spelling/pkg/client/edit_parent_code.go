package client

type EditParentCodeRequest struct {
	UserId     string `json:user_id`
	ParentCode string `json:"parent_code"`
}

type EditParentCodeResponse struct {
}
