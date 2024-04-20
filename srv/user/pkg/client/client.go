package client

import (
	"context"
	"github.com/google/uuid"
	http "github.com/sabrina-djebbar/spelling-app-backend/lib/shttp/client"
	"github.com/sabrina-djebbar/spelling-app-backend/srv/user/pkg/models"
)

const (
	GetUserPath           = "/get_user"
	CreateUserPath        = "/create_user"
	LoginPath             = "/login"
	logoutPath            = "/logout"
	editUserPath          = "/edit_user"
	editParentDetailsPath = "/edit_parent_details"
)

type Client interface {
	GetUser(req GetUserRequest) (*models.User, error)
	CreateUser(req CreateUserRequest) (*models.User, error)
	EditUser(req EditUserRequest) (*models.User, error)
	EditParentDetails(req EditParentDetailsRequest) (*models.User, error)
	Login(req LoginRequest) (*models.User, error)
}

type client struct {
	internal *http.InternalClient
}

func New() *client {
	cfg := http.InternalClientOptions{

		Host:    "http://user",
		Timeout: 5,
	}

	/* s := http.Server{
		Addr:         *bindAddress,      // configure the bind address
		Handler:      sm,                // set the default handler
		ErrorLog:     l,                 // set the logger for the server
		ReadTimeout:  5 * time.Second,   // max time to read request from the client
		WriteTimeout: 10 * time.Second,  // max time to write response to the client
		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
	} */

	return &client{internal: http.NewInternalClient(cfg)}
}

func (c *client) GetUser(ctx context.Context, req GetUserRequest) (*GetUserResponse, error) {
	res := &GetUserResponse{}
	return res, c.internal.Do(ctx, GetUserPath, req, res)

}

func (c *client) CreateUser(ctx context.Context, req CreateUserRequest) (*CreateUserResponse, error) {
	res := &CreateUserResponse{}
	return res, c.internal.Do(ctx, CreateUserPath, req, res)
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
