package client

import "github.com/sabrina-djebbar/spelling-app-backend/srv/spelling/pkg/models"

type CreateSpellingSetRequest struct {
	Name           string   `json:"name"`
	RecommendedAge int      `json:"recommended_age,omitempty"`
	Description    string   `json:"description,omitempty"`
	Words          []string `json:"words"`
	Tags           string   `json:"tags"`
	Creator        string   `json:"creator,omitempty"`
}

type CreateSpellingSetResponse struct {
	SpellingSet *models.SpellingSet `json:"spelling_set"`
}
