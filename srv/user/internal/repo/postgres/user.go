package postgres

import (
	"database/sql"
	"errors"
	dbCli "user/internal/database/postgres"
	"user/pkg/client"
	"user/pkg/models"
)

type User = models.User
type UserService = client.UserService

// CreateUser create a User
func (user *User) CreateUser() (User, error) {
	sqlStatement := `
        INSERT INTO users (id, name, role, username, password, birthday, academic_year, class_id, parent_name, parent_email, parent_code, school_name, classes)
        VALUES ($1, $2, $3, $4,crypt($5,gen_salt('sha1')), $6, $7, $8, $9, $10, $11, $12, $13)
    `
	row := dbCli.DbClient.QueryRow(sqlStatement)

	switch err := row.Scan(&c); err {
	// success, no errors
	case nil:
		return c, nil
	// In case of no output rows
	case sql.ErrNoRows:
		return c, errors.New("Insert Failed")
	// other errors occured
	default:
		return c, err
	}
}
