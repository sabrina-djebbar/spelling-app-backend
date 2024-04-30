package rpc

import (
	"context"
	"github.com/sabrina-djebbar/spelling-app-backend/srv/spelling/internal/app"
	"github.com/sabrina-djebbar/spelling-app-backend/srv/spelling/pkg/client"
)

type RPC interface {
	CreateSpellingWord(ctx context.Context, req client.CreateSpellingWordRequest) (*client.CreateSpellingWordResponse, error)
	CreateSpellingSet(ctx context.Context, req client.CreateSpellingSetRequest) (*client.CreateSpellingSetResponse, error)
	ListSpellingSets(ctx context.Context, req client.ListSpellingSetsRequest) (*client.ListSpellingSetResponse, error)
	CreateSpellingAttempt(ctx context.Context, req client.CreateSpellingAttemptRequest) (*client.CreateSpellingAttemptResponse, error)
	ListSpellingExerciseByUser(ctx context.Context, req client.ListSpellingExercisesByUserRequest) (*client.ListSpellingExercisesByUserResponse, error)
}

type rpc struct {
	app app.App
}

func (r rpc) ListSpellingExerciseByUser(ctx context.Context, req client.ListSpellingExercisesByUserRequest) (*client.ListSpellingExercisesByUserResponse, error) {
	res, err := r.app.ListSpellingExercisesByUser(ctx, req)
	if err != nil {
		return nil, err
	}
	return &client.ListSpellingExercisesByUserResponse{Exercises: res}, nil
}

func New(app app.App) RPC {
	return &rpc{
		app,
	}
}
