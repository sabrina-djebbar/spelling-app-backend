package client

import "github.com/sabrina-djebbar/spelling-app-backend/srv/spelling/pkg/models"

type ListSpellingExercisesByUserRequest struct {
	UserID string `json:"user_id"`
}

type ListSpellingExercisesByUserResponse struct {
	Exercises []models.SpellingExercise `json:"exercises"`
}
