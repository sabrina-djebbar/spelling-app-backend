// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package repo

import (
	"database/sql"
)

type Credential struct {
	ID       string
	UserID   string
	Password string
}

type User struct {
	ID          string
	Username    string
	ParentCode  string
	DateOfBirth sql.NullTime
	Created     sql.NullTime
}