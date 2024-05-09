package rpc

import (
	"context"
	"github.com/sabrina-djebbar/spelling-app-backend/srv/user/pkg/client"
)

func (r rpc) ListUser(ctx context.Context, req client.ListUsersRequest) (*client.ListUsersResponse, error) {
	res, err := r.app.ListUsers(ctx)
	if err != nil {
		return nil, err
	}
	return &client.ListUsersResponse{Users: res}, nil
}
