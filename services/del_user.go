package services

import (
	"fmt"
	"github.com/kuzeofficial/golang-test-api/repositories"
)

func DelUser(key string) error {
	db, err := repositories.ConnectDatabase()
	if err != nil {
		return err
	}
	errorDel := db.Delete(key)
	if errorDel != nil {
		fmt.Println("failed to delete item:", errorDel)
		return errorDel
	}
	return errorDel
}
