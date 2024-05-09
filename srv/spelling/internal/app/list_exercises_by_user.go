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

	return TransformSpellingExercises(exercises), nil

}
