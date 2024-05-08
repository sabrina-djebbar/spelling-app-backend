package client

import (
	"github.com/sabrina-djebbar/spelling-app-backend/srv/spelling/pkg/models"
	"time"
)

type CreateSpellingAttemptRequest struct {
	AttemptID     string    `json:"attempt_id,omitempty"`
	UserID        string    `json:"user_id"`
	SetID         string    `json:"set_id"`
	WordID        string    `json:"word_id"`
	Spelling      string    `json:"spelling"`
	Score         float64   `json:"score"`
	NumOfAttempts int       `json:"num_of_attempts"`
	LastAttempt   time.Time `json:"last_attempt"`
}

type CreateSpellingAttemptResponse struct {
	Attempt *models.SpellingExercise `json:"attempt"`
}
