package app

import (
	"context"
	"fmt"
	"github.com/sabrina-djebbar/spelling-app-backend/lib/database"
	"github.com/sabrina-djebbar/spelling-app-backend/lib/id"
	"github.com/sabrina-djebbar/spelling-app-backend/srv/spelling/internal/infrastructure/repo"
	"github.com/sabrina-djebbar/spelling-app-backend/srv/spelling/pkg/client"
	"github.com/sabrina-djebbar/spelling-app-backend/srv/spelling/pkg/models"
)

func (a *app) CreateSpellingAttempt(ctx context.Context, req client.CreateSpellingAttemptRequest) (*models.SpellingExercise, error) {
	if req.AttemptID == "" {
		fmt.Println("AttemptID " + req.AttemptID + " will be generated")
		req.AttemptID = id.Generate("exercise")
	}
	attempts, err := a.repository.AddSpellingAttempt(ctx, repo.AddSpellingAttemptParams{ID: req.AttemptID,
		UserID:        req.UserID,
		SetID:         req.SetID,
		WordID:        req.WordID,
		Spelling:      req.Spelling,
		Score:         req.Score,
		NumOfAttempts: database.IntToInt32(req.NumOfAttempts),
		LastAttempt:   req.LastAttempt,
	})
	if err != nil {
		return nil, err
	}
	exercise := TransformSpellingExercises(attempts)
	return &exercise[0], nil
}
