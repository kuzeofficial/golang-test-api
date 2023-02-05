package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/kuzeofficial/golang-test-api/handlers"
)

func main() {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))
	app.Get("/users", handlers.GetUsersHandler)
	app.Get("/user/:key", handlers.GetUser)
	app.Post("/user/new", handlers.SetUser)
	app.Delete("/user/:key", handlers.DelUser)
	app.Listen(":3000")
}
