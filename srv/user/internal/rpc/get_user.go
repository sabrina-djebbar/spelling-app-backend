package rpc

import (
	"context"
	"github.com/sabrina-djebbar/spelling-app-backend/srv/user/pkg/client"
)

func (r rpc) GetUser(ctx context.Context, req client.GetUserRequest) (*client.GetUserResponse, error) {
	res, err := r.app.GetUser(ctx, req.UserID)
	if err != nil {
		return nil, err
	}
	return &client.GetUserResponse{User: res}, nil
}
