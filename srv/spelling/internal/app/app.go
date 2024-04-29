package app

import (
	"context"
	repo "github.com/sabrina-djebbar/spelling-app-backend/srv/spelling/internal/infrastructure"
	"github.com/sabrina-djebbar/spelling-app-backend/srv/spelling/pkg/client"
	"github.com/sabrina-djebbar/spelling-app-backend/srv/spelling/pkg/models"
	"strings"
)

type App interface {
	CreateSpellingWord(ctx context.Context, req client.CreateSpellingWordRequest) (*models.SpellingWord, error)
	CreateSpellingSet(ctx context.Context, req client.CreateSpellingSetRequest) (*models.SpellingSet, error)
	ListSpellingSets(ctx context.Context, req client.ListSpellingSetsRequest) ([]models.SpellingSet, error)
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

func FormatTagsStringToArray(spelling string) []string {
	spelling = strings.ReplaceAll(spelling, " ", "")
	return strings.Split(spelling, ",")
}
