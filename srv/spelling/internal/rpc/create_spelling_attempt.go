package rpc

import (
	"context"
	"github.com/sabrina-djebbar/spelling-app-backend/srv/spelling/pkg/client"
)

func (r rpc) CreateSpellingAttempt(ctx context.Context, req client.CreateSpellingAttemptRequest) (*client.CreateSpellingAttemptResponse, error) {
	res, err := r.app.CreateSpellingAttempt(ctx, req)
	if err != nil {
		return nil, err
	}
	return &client.CreateSpellingAttemptResponse{Attempt: res}, nil
}
