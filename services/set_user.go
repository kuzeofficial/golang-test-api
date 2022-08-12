package services

import (
	"fmt"
	"github.com/deta/deta-go/service/base"
	model "github.com/kuzeofficial/golang-test-api/models"
	"github.com/kuzeofficial/golang-test-api/repositories"
	utilities "github.com/kuzeofficial/golang-test-api/utils"
)

func AddUser(user *model.User) error {
	variables := utilities.GetEnviromentsVariables()
	d := repositories.ConnectDatabase()
	db, err := base.New(d, variables.DetaBaseName)
	_, err = db.Insert(&model.User{
		Name:  user.Name,
		Title: user.Title,
	})
	if err != nil {
		fmt.Println("failed to insert item:", err)
		return err
	}
	return err
}
