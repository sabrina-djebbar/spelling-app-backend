package scriteria

import (
	"encoding/csv"
	"fmt"
	"github.com/sabrina-djebbar/spelling-app-backend/lib/database"
	"github.com/sabrina-djebbar/spelling-app-backend/lib/serr"
	"os"
)

type Pattern struct {
	Difficulty         float64
	YearGroup          string
	PatternDescription string
	Regex              string
}

func readPatternsFromFile(filename string) ([]Pattern, error) {
	// Open the CSV file
	file, err := os.Open(filename)
	if err != nil {
		return nil, serr.Wrap(err, serr.WithMessage("cannot open csv file, "+filename+"\n"))
	}
	defer file.Close()

	// Create a CSV reader
	reader := csv.NewReader(file)

	// Read all the records from CSV
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	// Initialize an empty array to store patterns
	var patterns []Pattern

	// Iterate over each record and create a Pattern struct
	for _, record := range records {
		pattern := Pattern{
			Difficulty:         database.StringToFloat64(record[1]),
			YearGroup:          record[2],
			PatternDescription: record[3],
			Regex:              record[4],
		}
		// Append the pattern to the array
		patterns = append(patterns, pattern)
	}

	return patterns, nil
}

func GetSpellingPattern() ([]Pattern, error) {
	filename := "spelling-app-backend/lib/scriteria/spelling_criteria.csv"
	patterns, err := readPatternsFromFile(filename)
	if err != nil {
		return nil, serr.Wrap(err, serr.WithMessage("unable to read spelling criteria"))
	}

	// Print out the patterns
	for _, pattern := range patterns {
		fmt.Printf("%+v\n", pattern)
	}
	return patterns, nil
}
