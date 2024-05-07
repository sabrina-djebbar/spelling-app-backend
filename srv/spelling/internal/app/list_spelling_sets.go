package app

import (
	"context"
	"github.com/sabrina-djebbar/spelling-app-backend/srv/spelling/pkg/client"
	"github.com/sabrina-djebbar/spelling-app-backend/srv/spelling/pkg/models"
)

func (a *app) ListSpellingSets(ctx context.Context, req client.ListSpellingSetsRequest) ([]models.SpellingSet, error) {
	sets, err := a.repository.ListSetsByTags(ctx, FormatTagsStringToArray(req.Tags))
	if err != nil {
		return nil, err
	}
	spellingSets := make([]models.SpellingSet, 0)
	setMap := make(map[string]int)
	for _, set := range sets {
		index, ok := setMap[set.ID]
		if ok {
			// If set ID is found in map, append the attempt to existing SpellingExercise
			spellingSets[index].Words = append(spellingSets[index].Words, set.Word)
		} else {
			// If exercise ID is not found in map, create a new SpellingExercise
			setMap[set.ID] = len(spellingSets)
			spellingSets = append(
				spellingSets,
				models.SpellingSet{
					ID:             set.ID,
					Name:           set.Name,
					RecommendedAge: set.RecommendedAge,
					Description:    set.Description,
					Tags:           FormatTagsStringToArray(set.Tags),
					Words: []models.SpellingWord{
						set.Word,
					},
				},
			)

		}
	}

	return spellingSets, nil
}
