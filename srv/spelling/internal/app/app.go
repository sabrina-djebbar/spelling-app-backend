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
	ListSpellingExercisesByUser(ctx context.Context, req client.ListSpellingExercisesByUserRequest) ([]models.SpellingExercise, error)
	CreateSpellingAttempt(ctx context.Context, req client.CreateSpellingAttemptRequest) (*models.SpellingExercise, error)
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

func TransformSpellingExercises(exercises []repo.SpellingExercise) []models.SpellingExercise {
	spellingExercises := make([]models.SpellingExercise, 0)
	exerciseMap := make(map[string]int)
	for _, exercise := range exercises {
		index, ok := exerciseMap[exercise.ID]
		if ok {
			// If exercise ID is found in map, append the attempt to existing SpellingExercise
			spellingExercises[index].Word = append(spellingExercises[index].Word, exercise.Attempt)
		} else {
			// If exercise ID is not found in map, create a new SpellingExercise
			exerciseMap[exercise.ID] = len(spellingExercises)
			spellingExercises = append(
				spellingExercises,
				models.SpellingExercise{
					ID:     exercise.ID,
					UserID: exercise.UserID,
					Set:    exercise.Set,
					Word: []models.SpellingAttempt{
						exercise.Attempt,
					},
				},
			)

		}
	}
	return spellingExercises
}
