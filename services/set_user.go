package services

import (
	"fmt"

	model "github.com/kuzeofficial/golang-test-api/models"
	"github.com/kuzeofficial/golang-test-api/repositories"
)

func AddUser(user *model.User) error {
	db, err := repositories.ConnectDatabase()
	_, err = db.Insert(&model.User{
		Name:  user.Name,
		Email: user.Email,
		Permission: user.Permission,
	})
	if err != nil {
		fmt.Println("failed to insert item:", err)
		return err
	}
	return err
}
