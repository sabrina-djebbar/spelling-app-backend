package rpc

import (
	"context"
	"github.com/sabrina-djebbar/spelling-app-backend/srv/user/pkg/client"
)

func (r rpc) EditUser(ctx context.Context, req client.EditUserRequest) (*client.EditUserResponse, error) {
	u, err := r.app.EditUser(ctx, req)
	if err != nil {
		return nil, err
	}
	return &client.EditUserResponse{User: u}, nil
}
