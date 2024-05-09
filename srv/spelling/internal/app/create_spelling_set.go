package app

import (
	"context"
	"github.com/sabrina-djebbar/spelling-app-backend/lib/database"
	"github.com/sabrina-djebbar/spelling-app-backend/lib/serr"
	repo "github.com/sabrina-djebbar/spelling-app-backend/srv/spelling/internal/infrastructure"
	"github.com/sabrina-djebbar/spelling-app-backend/srv/spelling/pkg/client"
	"github.com/sabrina-djebbar/spelling-app-backend/srv/spelling/pkg/models"
)

func (a *app) CreateSpellingSet(ctx context.Context, req client.CreateSpellingSetRequest) (*models.SpellingSet, error) {
	// first create spelling set
	// for each word add word to list
	if req.RecommendedAge <= "" {
		age, err := a.CalculateRecommendedAgeGroup(ctx, req.Words)
		if err != nil {
			return nil, serr.Wrap(err, serr.WithMessage("unable to calculate recommended age"))
		}
		req.RecommendedAge = age
	}

	repoSet, err := a.repository.CreateSpellingSet(ctx, repo.CreateSpellingSetParams{Name: req.Name, Description: req.Description, RecommendedAge: req.RecommendedAge, Creator: req.Creator, Tags: req.Tags})
	if err != nil {
		return nil, err
	}

	set := &models.SpellingSet{
		ID:             repoSet.ID,
		Name:           repoSet.Name,
		RecommendedAge: repoSet.RecommendedAge,
		Description:    database.SQLNullStringToString(repoSet.Description),
		Words:          nil,
		// TODO: add tags based on words, can be done on the FE though
		Tags: FormatTagsStringToArray(database.SQLNullStringToString(repoSet.Tags)),
	}

	for _, w := range req.Words {
		word, err := a.repository.AddWordToSet(ctx, set.ID, w)
		if err != nil {
			return nil, err
		}
		// push word to set
		set.Words = append(set.Words, models.SpellingWord{
			ID:                   word.ID,
			Spelling:             word.Spelling,
			Definition:           database.SQLNullStringToString(word.Definition),
			Difficulty:           word.Difficulty,
			TotalAvailablePoints: database.Int32ToInt(word.TotalAvailablePoints),
			Created:              database.SQLNullTimeToTime(word.Created),
			Tags:                 FormatTagsStringToArray(database.SQLNullStringToString(word.Tags)),
		})
	}
	return set, nil
}

func (a *app) CalculateRecommendedAgeGroup(ctx context.Context, words []string) (string, error) {
	var difficulty = 0.00
	for _, word := range words {
		wordDifficulty, err := a.repository.GetSpellingWordDifficulty(ctx, word)
		if err != nil {
			return "", err
		}
		difficulty += wordDifficulty
	}
	difficulty = difficulty / float64(len(words))
	recommendedAge := FindAgeForDifficulty(difficulty)
	return recommendedAge, nil
}

// TODO: make recommended_age a string and use year group instead
func FindAgeForDifficulty(difficulty float64) string {
	switch {
	case difficulty < 10.00:
		return "reception"
	case difficulty >= 10.00 && difficulty < 20.00:
		return "Year 1"
	case difficulty >= 20.00 && difficulty < 30.00:
		return "Year 2"
	default:
		return "N/A"
	}
}
