package rpc

import (
	"context"
	"github.com/sabrina-djebbar/spelling-app-backend/srv/spelling/pkg/client"
)

func (r rpc) CreateSpellingWord(ctx context.Context, req client.CreateSpellingWordRequest) (*client.CreateSpellingWordResponse, error) {
	word, err := r.app.CreateSpellingWord(ctx, req)
	if err != nil {
		return nil, err
	}
	return &client.CreateSpellingWordResponse{Word: word}, nil
}
