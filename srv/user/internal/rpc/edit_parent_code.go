package rpc

import (
	"context"
	"github.com/sabrina-djebbar/spelling-app-backend/srv/user/pkg/client"
)

func (r rpc) EditParentCode(ctx context.Context, req client.EditParentCodeRequest) (*client.EditParentCodeResponse, error) {
	u, err := r.app.EditParentDetails(ctx, req)
	if err != nil {
		return nil, err
	}
	return &client.EditParentCodeResponse{User: *u}, nil
}
