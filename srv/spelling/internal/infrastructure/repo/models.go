// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package repo

import (
	"database/sql"
	"time"
)

type SpellingExercise struct {
	ID            string
	UserID        string
	SetID         string
	WordID        string
	Score         float64
	NumOfAttempts int32
	LastAttempt   time.Time
}

type SpellingSet struct {
	ID             string
	Name           string
	RecommendedAge string
	Description    sql.NullString
	Tags           sql.NullString
	Creator        string
	Created        sql.NullTime
}

type SpellingSetWord struct {
	SetID  string
	WordID string
}

type SpellingWord struct {
	ID                   string
	Spelling             string
	Definition           sql.NullString
	Class                string
	Tags                 sql.NullString
	Difficulty           float64
	TotalAvailablePoints int32
	Created              sql.NullTime
}
