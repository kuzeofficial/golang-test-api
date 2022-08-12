package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
    app := fiber.New()
		app.Use(cors.New(cors.Config{
			AllowOrigins: "*",
		}))
    // app.Get("/", helloWorldHandler)
		app.Get("/users", userHandler)
		app.Static("/", "./client/dist")
    app.Listen(":3000")
}

func userHandler(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"name": "John Doe",
			"age":  42,
		})
}

// func helloWorldHandler(c *fiber.Ctx) error {
// 	return c.SendString("Hello, World 👋!")
// }