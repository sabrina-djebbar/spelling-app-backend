package app

import (
	"context"
	"github.com/sabrina-djebbar/spelling-app-backend/srv/spelling/pkg/client"
	"github.com/sabrina-djebbar/spelling-app-backend/srv/spelling/pkg/models"
)

func (a *app) ListSpellingExercisesByUser(ctx context.Context, req client.ListSpellingExercisesByUserRequest) ([]models.SpellingExercise, error) {
	exercises, err := a.repository.ListExercisesByUserID(ctx, req.UserID)
	if err != nil {
		return nil, err
	}
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

	return spellingExercises, nil

}
