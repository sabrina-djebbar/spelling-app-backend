package rpc

import (
	"context"
	"github.com/sabrina-djebbar/spelling-app-backend/srv/spelling/pkg/client"
)

func (r rpc) CreateSpellingSet(ctx context.Context, req client.CreateSpellingSetRequest) (*client.CreateSpellingSetResponse, error) {
	if req.Creator == "" {
		req.Creator = "spellatrix"
	}
	set, err := r.app.CreateSpellingSet(ctx, req)
	if err != nil {
		return nil, err
	}
	return &client.CreateSpellingSetResponse{SpellingSet: set}, nil
}
