package services

import (
	"fmt"
	"github.com/deta/deta-go/service/base"
	model "github.com/kuzeofficial/golang-test-api/models"
	"github.com/kuzeofficial/golang-test-api/repositories"
)

func GetUser(key string) []*model.User {
	db, err := repositories.ConnectDatabase()
	var users []*model.User
	_, err = db.Fetch(&base.FetchInput{
		Q: base.Query{
			{"key": key},
		},
		Dest: &users,
	})
	if err != nil {
		fmt.Println("failed to fetch items:", err)
		return users
	}
	return users
}
