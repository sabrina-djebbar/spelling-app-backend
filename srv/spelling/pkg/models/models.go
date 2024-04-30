package models

import "time"

type Class string

const (
	Noun        Class = "Noun"
	Verb        Class = "Verb"
	Adverb      Class = "Adverb"
	Conjuction  Class = "Conjuction"
	Preposition Class = "Preposition"
	Adjective   Class = "Adjective"
	Pronoun     Class = "Pronoun"
	Determiner  Class = "Determiner"
)

// SpellingWord represents the spelling_word table
type SpellingWord struct {
	ID                   string    `json:"id"`
	Spelling             string    `json:"spelling"`
	Definition           string    `json:"definition"`
	Class                Class     `json:"class"`
	Difficulty           float64   `json:"difficulty"`
	TotalAvailablePoints int       `json:"total_available_points"`
	Tags                 []string  `json:"tags"`
	Created              time.Time `json:"created"`
}

// SpellingSet represents the spelling_set table
type SpellingSet struct {
	ID             string         `json:"id"`
	Name           string         `json:"name"`
	RecommendedAge string         `json:"recommended_age"`
	Description    string         `json:"description"`
	Tags           []string       `json:"tags"`
	Words          []SpellingWord `json:"words"`
}

// SpellingExercise represents the spelling_exercise table
type SpellingExerciseV0 struct {
	ID            string    `json:"id"`
	UserID        string    `json:"user_id"`
	SetID         string    `json:"set_id"`
	WordID        string    `json:"word_id"`
	Spelling      string    `json:"spelling"`
	Score         float64   `json:"score"`
	NumOfAttempts int       `json:"num_of_attempts"`
	LastAttempt   time.Time `json:"last_attempt"`
}

type SpellingExercise struct {
	ID     string            `json:"id"`
	UserID string            `json:"user_id"`
	Set    SpellingSet       `json:"set"`
	Word   []SpellingAttempt `json:"word_attempts"`
}

type SpellingAttempt struct {
	Word          SpellingWord `json:"word"`
	Spelling      string       `json:"spelling"`
	Score         float64      `json:"score"`
	NumOfAttempts int          `json:"num_of_attempts"`
	LastAttempt   time.Time    `json:"last_attempt"`
}
