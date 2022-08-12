package main

import (
	"fmt"

	"github.com/deta/deta-go/service/base"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/kuzeofficial/golang-test-api/handlers"
	model "github.com/kuzeofficial/golang-test-api/models"
	"github.com/kuzeofficial/golang-test-api/repositories"
	utilities "github.com/kuzeofficial/golang-test-api/utils"
)







func main() {
    app := fiber.New()
		app.Use(cors.New(cors.Config{
			AllowOrigins: "*",
		}))
    // app.Get("/", helloWorldHandler)
		app.Get("/users", handlers.GetUsersHandler)
		app.Get("/user/:id", handlers.GetUser)
		app.Post("/user/new", AddUser)
		app.Static("/", "./client/dist")
    app.Listen(":3000")
}
// func userPutHandler(c *fiber.Ctx) error {
// 		var user User
// 		 err := json.Unmarshal(c.Body(), &user)
// 		 if err == nil {
// 			fmt.Printf("%+v\n", user)
// 			return c.JSON(fiber.Map{
// 				"message": "Usuario Creado con exito!",
// 			})
// 		 }
// 		 return c.Status(500).JSON(fiber.Map{
// 			"message": "Error al crear al usuario",
// 		 })


// }
// func userHandler(c *fiber.Ctx) error {
// 		return c.JSON(fiber.Map{
// 			"name": "John Doe",
// 			"age":  42,
// 		})
// }

// func getUser(c *fiber.Ctx) error {
// 	fmt.Print(c.AllParams()) // "{"name": "fenny"}"
// 	return c.JSON(fiber.Map{
// 		"name": "John Doe",
// 	})
// 	// ...
// }


func AddUser(c *fiber.Ctx) error {
	//Connect to database
	d := repositories.ConnectDatabase()

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
	db, err := base.New(d, "base_name")
	_, err = db.Insert(&model.User{
		Name:  user.Name,
		Title: user.Title,
	})
	if err != nil {
			fmt.Println("failed to insert item:", err)
			return c.Status(fiber.StatusBadRequest).JSON(errors)
	}
	return c.JSON(fiber.Map{
		"message": "Usuario Creado con exito!",
	})
}