package client

import "time"

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
	ID string `json:"id"`
}
