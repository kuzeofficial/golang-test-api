package utilities

import (
	"github.com/joho/godotenv"
	model "github.com/kuzeofficial/golang-test-api/models"
	"log"
	"os"
)

func GetEnviromentsVariables() model.Env {
	err := godotenv.Load()
	var variables model.Env
	if err != nil {
		log.Fatalln("Error al cargar las variables")
	}
	variables.DetaProjectKey = os.Getenv("DETA_PROJECT_KEY")
	variables.DetaBaseName = os.Getenv("DETA_BASE_NAME")
	return variables
}
