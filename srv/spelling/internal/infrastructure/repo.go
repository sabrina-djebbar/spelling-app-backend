package infrastructure

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"github.com/sabrina-djebbar/spelling-app-backend/lib/database"
	"github.com/sabrina-djebbar/spelling-app-backend/lib/id"
	"github.com/sabrina-djebbar/spelling-app-backend/lib/serr"
	"github.com/sabrina-djebbar/spelling-app-backend/srv/spelling/internal/infrastructure/repo"
	"github.com/sabrina-djebbar/spelling-app-backend/srv/spelling/pkg/models"
)

type Repository struct {
	q repo.Queries
}

func NewRepo(queries repo.Queries) Repository {
	return Repository{q: queries}
}

type CreateSpellingWordParams struct {
	Spelling             string
	Definition           string
	Difficulty           float64
	TotalAvailablePoints int
	Class                models.Class
}

func (r *Repository) CreateSpellingWord(ctx context.Context, args CreateSpellingWordParams) (*repo.SpellingWord, error) {
	params := repo.CreateSpellingWordParams{
		ID:                   id.Generate("word"),
		Spelling:             args.Spelling,
		Class:                string(args.Class),
		Difficulty:           args.Difficulty,
		Definition:           database.StringToSQLNullString(args.Definition),
		TotalAvailablePoints: database.IntToInt32(args.TotalAvailablePoints),
	}

	word, err := r.q.CreateSpellingWord(ctx, params)
	if err != nil {
		return nil, serr.Wrap(err, serr.WithMessage("Unable to create word"))
	}

	return &word, nil
}

type CreateSpellingSetParams struct {
	Name           string
	Description    string
	RecommendedAge int
	Creator        string
}

func (r *Repository) CreateSpellingSet(ctx context.Context, args CreateSpellingSetParams) (*repo.SpellingSet, error) {
	params := repo.CreateSpellingSetParams{
		ID:             id.Generate("set"),
		Name:           args.Name,
		Description:    database.StringToSQLNullString(args.Description),
		RecommendedAge: database.IntToInt32(args.RecommendedAge),
		Creator:        args.Creator,
	}

	set, err := r.q.CreateSpellingSet(ctx, params)
	if err != nil {
		return nil, serr.Wrap(err, serr.WithMessage("Unable to create scriteria set"))
	}
	return &set, nil
}

func (r *Repository) AddWordToSet(ctx context.Context, setID string, wordID string) (*repo.SpellingWord, error) {
	word, err := r.q.GetSpellingWord(ctx, wordID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, serr.Wrap(err, serr.WithMessage("No word found with id "+wordID))
		}
		return nil, serr.Wrap(err, serr.WithMessage("Unable to find word"))
	}
	err = r.q.AddWordToSet(ctx,
		repo.AddWordToSetParams{
			SetID:  setID,
			WordID: wordID,
		})
	if err != nil {
		return nil, serr.Wrap(err, serr.WithMessage("Unable to add word to set"))
	}
	return &word, nil
}

func (r *Repository) GetSpellingWordDifficulty(ctx context.Context, wordID string) (float64, error) {
	difficulty, err := r.q.GetWordDifficulty(ctx, wordID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return 0, serr.Wrap(err, serr.WithMessage("No word found with id "+wordID))
		}
		return 0, serr.Wrap(err, serr.WithMessage("Unable to get word difficulty"))
	}
	return difficulty, nil
}
