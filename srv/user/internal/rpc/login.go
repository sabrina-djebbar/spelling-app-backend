package rpc

import (
	"context"
	"github.com/sabrina-djebbar/spelling-app-backend/srv/user/pkg/client"
)

func (r rpc) Login(ctx context.Context, req client.LoginRequest) (*client.LoginResponse, error) {

	res, err := r.app.Login(ctx, req)
	if err != nil {
		return nil, err
	}
	return &client.LoginResponse{User: res}, nil
}
