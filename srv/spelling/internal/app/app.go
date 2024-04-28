package app

import (
	"context"
	repo "github.com/sabrina-djebbar/spelling-app-backend/srv/spelling/internal/infrastructure"
	"github.com/sabrina-djebbar/spelling-app-backend/srv/spelling/pkg/client"
	"github.com/sabrina-djebbar/spelling-app-backend/srv/spelling/pkg/models"
)

type App interface {
	CreateSpellingWord(ctx context.Context, req client.CreateSpellingWordRequest) (*models.SpellingWord, error)
	CreateSpellingSet(ctx context.Context, req client.CreateSpellingSetRequest) (*models.SpellingSet, error)
}

type app struct {
	repository repo.Repository
}

func New(
	repository repo.Repository,
) App {
	return &app{
		repository,
	}
}
