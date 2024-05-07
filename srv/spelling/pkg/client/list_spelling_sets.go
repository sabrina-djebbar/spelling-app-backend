package client

import "github.com/sabrina-djebbar/spelling-app-backend/srv/spelling/pkg/models"

type ListSpellingSetsRequest struct {
	Tags string `json:"tags;omitempty"`
}

type ListSpellingSetResponse struct {
	Sets []models.SpellingSet `json:"sets"`
}
