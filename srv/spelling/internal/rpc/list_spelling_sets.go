package rpc

import (
	"context"
	"github.com/sabrina-djebbar/spelling-app-backend/srv/spelling/pkg/client"
)

func (r rpc) ListSpellingSets(ctx context.Context, req client.ListSpellingSetsRequest) (*client.ListSpellingSetResponse, error) {
	sets, err := r.app.ListSpellingSets(ctx, req)
	if err != nil {
		return nil, err
	}
	return &client.ListSpellingSetResponse{Sets: sets}, nil
}
