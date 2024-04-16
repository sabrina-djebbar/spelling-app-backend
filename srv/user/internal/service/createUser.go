package service
import ("fmt"
 db "internal/repo/postgres")

 var srv = db.UserService;
 
 func CreateUser(u *db.UCreateUser) error {
	fmt.Println("In service")
	us = u
	err = us.CreateUser()
	if err != nil {
		return err
	}
	return nil
}