package repositories

import (
	"fmt"
	utilities "github.com/kuzeofficial/golang-test-api/utils"

	"github.com/deta/deta-go/deta"
)

func ConnectDatabase() *deta.Deta {
	variables := utilities.GetEnviromentsVariables()
	d, err := deta.New(deta.WithProjectKey(variables.DetaProjectKey))
	if err != nil {
		fmt.Println("failed to init new Deta instance:", err)
		return d
	}
	return d
}
