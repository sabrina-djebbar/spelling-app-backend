package rpc

import (
	"context"
	"github.com/sabrina-djebbar/spelling-app-backend/srv/user/pkg/client"
)

func (r rpc) CreateUser(ctx context.Context, req client.CreateUserRequest) (*client.CreateUserResponse, error) {
	res, err := r.app.CreateUser(ctx, req)
	if err != nil {
		return nil, err
	}
	return &client.CreateUserResponse{User: res}, nil
}
