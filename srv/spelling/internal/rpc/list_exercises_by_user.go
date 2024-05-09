package rpc

import (
	"context"
	"github.com/sabrina-djebbar/spelling-app-backend/srv/spelling/pkg/client"
)

func (r rpc) ListSpellingExerciseByUser(ctx context.Context, req client.ListSpellingExercisesByUserRequest) (*client.ListSpellingExercisesByUserResponse, error) {
	res, err := r.app.ListSpellingExercisesByUser(ctx, req)
	if err != nil {
		return nil, err
	}
	return &client.ListSpellingExercisesByUserResponse{Exercises: res}, nil
}
