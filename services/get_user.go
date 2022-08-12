package services

import (
	"fmt"

	"github.com/deta/deta-go/service/base"
	model "github.com/kuzeofficial/golang-test-api/models"
	"github.com/kuzeofficial/golang-test-api/repositories"
)

func GetUser(id string) []*model.User {
	d := repositories.ConnectDatabase()
	db, err := base.New(d, "base_name")
	var users []*model.User
	_, err = db.Fetch(&base.FetchInput{
		Q: base.Query{
			{"id": id},
		},
		Dest: &users,
	})
	if err != nil {
		fmt.Println("failed to fetch items:", err)
		return users
	}
	return users
}