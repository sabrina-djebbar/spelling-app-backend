package postgres

import (
	"fmt"
	db "user/internal/database/postgres"
	"user/pkg/client"
)

type User = models.User
type UserService = client.UserService

//CreateUser create a User
func (user *User) CreateUser() error {
	fmt.Println("In Repo")

	stmt, err := dv.DbClient.Prepare("INSERT INTO public.User (User_name,category_id) VALUES ($1,$2);")
	if err != nil {
		return err
	}
	//closing the statement to prevent memory leaks
	defer stmt.Close()
	_, err = stmt.Exec(b.UserName, b.CategoryID)
	if err != nil {
		return err
	}
	return nil
}