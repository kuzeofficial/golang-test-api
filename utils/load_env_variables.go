package utilities

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	model "github.com/kuzeofficial/golang-test-api/models"
)

func GetEnviromentsVariables() model.Env {
	err := godotenv.Load("../.env")
	var variables model.Env
	if err != nil {
		log.Fatalln(err)
		log.Fatalln("Error al cargar las variables")
	}
	variables.DetaProjectKey = os.Getenv("DETA_PROJECT_KEY")
	variables.DetaBaseName = os.Getenv("DETA_BASE_NAME")
	return variables
}
