package repositories

import (
	"fmt"

	"github.com/deta/deta-go/deta"
)

func ConnectDatabase() *deta.Deta {
	d, err := deta.New(deta.WithProjectKey("b074ax0y_QmjZ3p9MmsJVbYp3d7vkGeeWNjxHU6yW"))
	if err != nil {
		fmt.Println("failed to init new Deta instance:", err)
		return d
	}
	return d
}