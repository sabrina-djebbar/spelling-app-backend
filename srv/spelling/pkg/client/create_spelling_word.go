package client

import "github.com/sabrina-djebbar/spelling-app-backend/srv/spelling/pkg/models"

type CreateSpellingWordRequest struct {
	Spelling   string `json:"spelling"`
	Class      string `json:"class"`
	Tags       string `json:"tags"`
	Definition string `json:"definition,omitempty"`
}

type CreateSpellingWordResponse struct {
	Word *models.SpellingWord `json:"word"`
}
