package handlers

import (
	"github.com/gofiber/fiber/v2"
	model "github.com/kuzeofficial/golang-test-api/models"
	"github.com/kuzeofficial/golang-test-api/services"
	utilities "github.com/kuzeofficial/golang-test-api/utils"
)

func SetUser(c *fiber.Ctx) error {
	user := new(model.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	errors := utilities.ValidateStruct(*user)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}
	err := services.AddUser(user)
	if err != nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"message": "The user has been exist!",
			"status":  "FAIL",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":     "User created correctly",
		"status_code": "OK",
	})
}
