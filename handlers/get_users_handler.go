package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kuzeofficial/golang-test-api/services"
)

func GetUsersHandler(c *fiber.Ctx) error {
	users := services.GetUsers()
	if len(users) == 0 {
		return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
			"message": "No hay usuarios",
		})
	}
	return c.JSON(users)
}