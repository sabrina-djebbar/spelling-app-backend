package app

import (
	"context"
	"github.com/sabrina-djebbar/spelling-app-backend/srv/spelling/pkg/client"
	"github.com/sabrina-djebbar/spelling-app-backend/srv/spelling/pkg/models"
)

func (a *app) ListSpellingSets(ctx context.Context, req client.ListSpellingSetsRequest) ([]models.SpellingSet, error) {
	sets, err := a.repository.ListSetsByTags(ctx, req.Tags)
	if err != nil {
		return nil, err
	}
	var spellingSets []models.SpellingSet
	for _, set := range sets {
		spellingSets = append(spellingSets, models.SpellingSet{
			ID:             set.ID,
			Name:           set.Name,
			RecommendedAge: set.RecommendedAge,
			Description:    set.Description,
			Tags:           FormatTagsStringToArray(set.Tags),
			Words:          set.Words,
		})
	}
	return spellingSets, nil
}
