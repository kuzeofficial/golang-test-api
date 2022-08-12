package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kuzeofficial/golang-test-api/services"
)

func GetUser(c *fiber.Ctx) error {
	param := c.AllParams()
	var paramId = ""
	for paramName, paramValue := range param {
		if paramName == "id" {
			paramId = paramValue
		}
	}
	users := services.GetUser(paramId)
	if len(users) == 0 {
		return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
			"message": "No hay usuarios",
		})
	}
	return c.JSON(users)
}