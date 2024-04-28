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
	// first create scriteria set
	// for each word add word to list
	if req.RecommendedAge <= 0 {
		age, err := a.CalculateRecommendedAge(ctx, req.Words)
		if err != nil {
			return nil, serr.Wrap(err, serr.WithMessage("unable to calculate recommended age"))
		}
		req.RecommendedAge = age
	}

	repoSet, err := a.repository.CreateSpellingSet(ctx, repo.CreateSpellingSetParams{Name: req.Name, Description: req.Description, RecommendedAge: req.RecommendedAge, Creator: req.Creator})
	if err != nil {
		return nil, err
	}

	set := &models.SpellingSet{
		ID:             repoSet.ID,
		Name:           repoSet.Name,
		RecommendedAge: database.Int32ToInt(repoSet.RecommendedAge),
		Description:    database.SQLNullStringToString(repoSet.Description),
		Words:          nil,
		// TODO: add tags based on words, can be done on the FE though
		Tags: database.SQLNullStringToString(repoSet.Tags),
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
		})
	}
	return set, nil
}

func (a *app) CalculateRecommendedAge(ctx context.Context, words []string) (int, error) {

	var difficulty = 0.00
	for _, word := range words {
		wordDifficulty, err := a.repository.GetSpellingWordDifficulty(ctx, word)
		if err != nil {
			return 0, err
		}
		difficulty += wordDifficulty
	}
	difficulty = difficulty / float64(len(words))
	recommendedAge := FindAgeForDifficulty(difficulty)
	return recommendedAge, nil
}

// TODO: make recommended_age a string and use year group instead
func FindAgeForDifficulty(difficulty float64) int {
	switch {
	case difficulty < 1000.00:
		return 0
	case difficulty >= 1000.00 && difficulty < 2000.00:
		return 1
	case difficulty >= 2000.00 && difficulty < 3000.00:
		return 2
	default:
		return 0
	}
}
