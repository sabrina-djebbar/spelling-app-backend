package models

import "time"

type User struct {
	ID         string    `json:"id"`
	Username   string    `json:"username"`
	Birthday   time.Time `json:"date_of_birth,omitempty"`
	ParentCode string    `json:parent_code`
}
