package rpc

import (
	"context"
	"github.com/sabrina-djebbar/spelling-app-backend/srv/user/internal/app"
	"github.com/sabrina-djebbar/spelling-app-backend/srv/user/pkg/client"
)

type RPC interface {
	GetUser(ctx context.Context, req client.GetUserRequest) (*client.GetUserResponse, error)
	CreateUser(ctx context.Context, req client.CreateUserRequest) (*client.CreateUserResponse, error)
	ListUser(ctx context.Context) (*client.ListUsersResponse, error)
}

type rpc struct {
	app app.App
}

func New(app app.App) RPC {
	return &rpc{
		app,
	}
}
