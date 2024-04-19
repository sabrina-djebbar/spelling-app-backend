// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: query.sql

package repo

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createCredentials = `-- name: CreateCredentials :exec
INSERT INTO credentials (user_id, password) VALUES($1,crypt($2,'crypt-des'))
`

type CreateCredentialsParams struct {
	UserID pgtype.Int4
	Crypt  string
}

func (q *Queries) CreateCredentials(ctx context.Context, arg CreateCredentialsParams) error {
	_, err := q.db.Exec(ctx, createCredentials, arg.UserID, arg.Crypt)
	return err
}

const createUser = `-- name: CreateUser :one
INSERT INTO users (username, date_of_birth, parent_code) VALUES ($1, $2, $3) RETURNING id, username, parent_code, date_of_birth, created
`

type CreateUserParams struct {
	Username    string
	DateOfBirth pgtype.Date
	ParentCode  pgtype.Text
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, createUser, arg.Username, arg.DateOfBirth, arg.ParentCode)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.ParentCode,
		&i.DateOfBirth,
		&i.Created,
	)
	return i, err
}