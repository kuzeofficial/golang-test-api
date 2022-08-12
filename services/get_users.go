package services

import (
	"fmt"
	"github.com/deta/deta-go/service/base"
	model "github.com/kuzeofficial/golang-test-api/models"
	"github.com/kuzeofficial/golang-test-api/repositories"
	utilities "github.com/kuzeofficial/golang-test-api/utils"
)

func GetUsers() []*model.User {
	d := repositories.ConnectDatabase()
	variables := utilities.GetEnviromentsVariables()
	db, err := base.New(d, variables.DetaBaseName)
	var users []*model.User
	_, err = db.Fetch(&base.FetchInput{
		Dest: &users,
	})
	if err != nil {
		fmt.Println("failed to fetch items:", err)
		return users
	}
	return users
}
