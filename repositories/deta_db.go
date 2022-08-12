package repositories

import (
	"github.com/deta/deta-go/deta"
	"github.com/deta/deta-go/service/base"
	utilities "github.com/kuzeofficial/golang-test-api/utils"
	"log"
)

func ConnectDatabase() (*base.Base, error) {
	variables := utilities.GetEnviromentsVariables()
	d, err := deta.New(deta.WithProjectKey(variables.DetaProjectKey))
	if err != nil {
		log.Fatalln("failed to init new Deta instance:", err)
	}
	db, err := base.New(d, variables.DetaBaseName)
	return db, err
}
