package client

const (
	getUserPath           = "/get_user"
	createUserPath        = "/create_user"
	loginPath             = "/login"
	logoutPath            = "/logout"
	editUserPath          = "/edit_user"
	editParentDetailsPath = "/edit_parent_details"
)

type UserService interface {
	GetUser(req GetUserRequest) (*User, error)
	CreateUser(req CreateUserRequest) (*User, error)
	EditUser(req EditUserRequest) (*User, error)
	EditParentDetails(req EditParentDetailsRequest) (*User, error)
	Login(req LoginRequest) (*User, error)
	Logout(req LogoutRequest) error
}
type Client interface {
	GetUser(req GetUserRequest) (*User, error)
	CreateUser(req CreateUserRequest) (*User, error)
	EditUser(req EditUserRequest) (*User, error)
	EditParentDetails(req EditParentDetailsRequest) (*User, error)
	Login(req LoginRequest) (*User, error)
	Logout(req LogoutRequest) error
}

type client struct {
	internal *http.InternalClient
}

func NewFromEnv() *client {
	cfg := http.InternalClientOptions{
		Name:    "user",
		Host:    "http://user",
		Timeout: 5,
	}

	// config.LoadConfigItem("USER_CLIENT", &cfg)

	return &client{
		internal: http.NewInternalClient(cfg),
	}
}

func (c *client) GetUser(ctx context, req GetUserRequest) (*GetUserResponse, error) {
	res := &GetUserResponse{}
	return res, c.internal.Do(ctx, "get_user", req, res)

}

type CreateUserRequest struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	ParentCode  string `json:"parent_code"`
	DateOfBirth string `json:"date_of_birth"`
}

type EditUserRequest struct {
	Username string `json:"username"`
	Birthday string `json:"date_of_birth"`
}

type EditParentDetailsRequest struct {
	user       uuid.UUID `json:"user,required"`
	ParentCode string    `json:"parent_code"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}
