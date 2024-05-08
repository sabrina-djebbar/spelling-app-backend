package infrastructure

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"github.com/sabrina-djebbar/spelling-app-backend/lib/database"
	"github.com/sabrina-djebbar/spelling-app-backend/lib/serr"
	"github.com/sabrina-djebbar/spelling-app-backend/srv/spelling/internal/infrastructure/repo"
	"github.com/sabrina-djebbar/spelling-app-backend/srv/spelling/pkg/models"
	"time"
)

func (r *Repository) AddSpellingAttempt(ctx context.Context, params repo.AddSpellingAttemptParams) (*SpellingExercise, error) {
	attempt, err := r.q.AddSpellingAttempt(ctx, params)
	if err != nil {
		return nil, serr.Wrap(err, serr.WithMessage("Unable to create spelling attempt"))
	}

	return r.transformSpellingAttempt(attempt), nil

}

type SpellingExercise struct {
	ID            string
	UserID        string
	Set           models.SpellingSet
	Attempt       models.SpellingAttempt
	Spelling      string
	Score         float64
	NumOfAttempts int
	LastAttempt   time.Time
}

func (r *Repository) ListExercisesByUserID(ctx context.Context, userID string) ([]SpellingExercise, error) {
	spellingExercises := make([]SpellingExercise, 0)
	attempts, err := r.q.GetSpellingExerciseAttemptByUser(ctx, userID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, serr.Wrap(err, serr.WithMessage("No exercises found for user "+userID))
		}
		return nil, serr.Wrap(err, serr.WithMessage("Unable to get exercises for user "+userID))
	}

	for _, attempt := range attempts {
		spellingExercises = append(spellingExercises, r.transformSpellingExercise(attempt))
	}
	return spellingExercises, nil
}

func (r *Repository) transformSpellingExercise(attempt repo.GetSpellingExerciseAttemptByUserRow) SpellingExercise {
	return SpellingExercise{
		ID:     attempt.ExerciseID,
		UserID: attempt.UserID,
		Set: models.SpellingSet{
			ID:             attempt.SetID,
			Name:           attempt.SetName,
			RecommendedAge: attempt.RecommendedAge,
			Description:    database.SQLNullStringToString(attempt.Description),
			Tags:           FormatNullStringTags(attempt.SetTags),
		},
		Attempt: models.SpellingAttempt{
			Spelling:      attempt.SpellingAttempt,
			Score:         attempt.Score,
			LastAttempt:   attempt.LastAttempt,
			NumOfAttempts: database.Int32ToInt(attempt.NumOfAttempts),
			Word: models.SpellingWord{
				ID:                   attempt.WordID,
				Spelling:             attempt.CorrectSpelling,
				Definition:           database.SQLNullStringToString(attempt.Definition),
				Class:                models.Class(attempt.WordClass),
				Difficulty:           attempt.Difficulty,
				TotalAvailablePoints: database.Int32ToInt(attempt.TotalAvailablePoints),
				Tags:                 FormatNullStringTags(attempt.WordTags),
			}},
	}
}

func (r *Repository) transformSpellingAttempt(attempt repo.AddSpellingAttemptRow) *SpellingExercise {
	return &SpellingExercise{
		ID:     attempt.ExerciseID,
		UserID: attempt.UserID,
		Set: models.SpellingSet{
			ID:             attempt.SetID,
			Name:           attempt.SetName,
			RecommendedAge: attempt.RecommendedAge,
			Description:    database.SQLNullStringToString(attempt.Description),
			Tags:           FormatNullStringTags(attempt.SetTags),
		},
		Attempt: models.SpellingAttempt{
			Spelling:      attempt.SpellingAttempt,
			Score:         attempt.Score,
			LastAttempt:   attempt.LastAttempt,
			NumOfAttempts: database.Int32ToInt(attempt.NumOfAttempts),
			Word: models.SpellingWord{
				ID:                   attempt.WordID,
				Spelling:             attempt.CorrectSpelling,
				Definition:           database.SQLNullStringToString(attempt.Definition),
				Class:                models.Class(attempt.WordClass),
				Difficulty:           attempt.Difficulty,
				TotalAvailablePoints: database.Int32ToInt(attempt.TotalAvailablePoints),
				Tags:                 FormatNullStringTags(attempt.WordTags),
			}},
	}
}
