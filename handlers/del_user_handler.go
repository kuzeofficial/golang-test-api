package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kuzeofficial/golang-test-api/services"
)

func DelUser(c *fiber.Ctx) error {
	param := c.AllParams()
	var paramKey = ""
	for paramName, paramValue := range param {
		if paramName == "key" {
			paramKey = paramValue
		}
	}
	err := services.DelUser(paramKey)
	if err != nil {
		return c.Status(fiber.StatusExpectationFailed).JSON(fiber.Map{
			"message": "The user not exist!",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User has been deleted correctly!",
		"status":  "OK",
	})
}
