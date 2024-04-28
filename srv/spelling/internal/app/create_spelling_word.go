package app

import (
	"context"
	"fmt"
	"github.com/sabrina-djebbar/spelling-app-backend/lib/database"
	"github.com/sabrina-djebbar/spelling-app-backend/lib/scriteria"
	repo "github.com/sabrina-djebbar/spelling-app-backend/srv/spelling/internal/infrastructure"
	"github.com/sabrina-djebbar/spelling-app-backend/srv/spelling/pkg/client"
	"github.com/sabrina-djebbar/spelling-app-backend/srv/spelling/pkg/models"
	"regexp"
)

func (a *app) CreateSpellingWord(ctx context.Context, req client.CreateSpellingWordRequest) (*models.SpellingWord, error) {
	difficulty, err := CalculateDifficulty(req.Spelling)
	if err != nil {
		return nil, err
	}
	args := repo.CreateSpellingWordParams{
		Spelling:             req.Spelling,
		Definition:           req.Definition,
		Class:                models.Class(req.Class),
		Difficulty:           difficulty,
		TotalAvailablePoints: len(req.Spelling) + 1,
	}
	word, err := a.repository.CreateSpellingWord(ctx, args)
	if err != nil {
		return nil, err
	}
	return &models.SpellingWord{
		ID:                   word.ID,
		Spelling:             word.Spelling,
		Definition:           database.SQLNullStringToString(word.Definition),
		Difficulty:           word.Difficulty,
		Class:                models.Class(word.Class),
		Tags:                 database.SQLNullStringToString(word.Tags),
		TotalAvailablePoints: database.Int32ToInt(word.TotalAvailablePoints),
	}, nil
}

func CalculateDifficulty(spelling string) (float64, error) {
	difficulty := float64(len(spelling))
	patterns := scriteria.GetSpellingPatterns()

	for _, pattern := range patterns {
		// Compile the regex pattern
		re, err := regexp.Compile(pattern.Regex)
		if err != nil {
			fmt.Println("Error compiling regex:", err)
			continue
		}
		// Check if the word matches the regex pattern
		if re.MatchString(spelling) {
			// Add the difficulty of the pattern to the total difficulty
			difficulty += pattern.Difficulty
		}
	}

	return difficulty, nil
}
