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
	EditUserPath          = "/edit_user"
	EditParentDetailsPath = "/edit_parent_details"
	ListUsersPath         = "/list_users"
)

type Client interface {
	GetUser(req GetUserRequest) (*models.User, error)
	CreateUser(req CreateUserRequest) (*models.User, error)
	EditUser(req EditUserRequest) (*models.User, error)
	EditParentDetails(req EditParentCodeRequest) (*models.User, error)
	Login(req LoginRequest) (*models.User, error)
	ListUsers(req ListUsersRequest) ([]*models.User, error)
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

func (c *client) ListUsers(ctx context.Context, req ListUsersRequest) (*ListUsersResponse, error) {
	res := &ListUsersResponse{}
	return res, c.internal.Do(ctx, ListUsersPath, req, res)
}

func (c *client) Login(ctx context.Context, req LoginRequest) (*LoginResponse, error) {
	res := &LoginResponse{}
	return res, c.internal.Do(ctx, LoginPath, req, res)
}
