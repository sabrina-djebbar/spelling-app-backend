package models

import "time"

type User struct {
	ID          string    `json:"id"`
	Username    string    `json:"username"`
	DateOfBirth time.Time `json:"date_of_birth"`
	ParentCode  string    `json:"parent_code"`
}
