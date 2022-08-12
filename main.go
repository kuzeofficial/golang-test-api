package main

import (
	"fmt"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/deta/deta-go/deta"
	"github.com/deta/deta-go/service/base"
)


type User struct {
	Key   string `json:"key"`
	Name  string `json:"name" validate:"required"`
	Title string `json:"title" validate:"required"`
}

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}
var validate = validator.New()
func ValidateStruct(user User) []*ErrorResponse {
	var errors []*ErrorResponse
	err := validate.Struct(user)
	if err != nil {
			for _, err := range err.(validator.ValidationErrors) {
					var element ErrorResponse
					element.FailedField = err.StructNamespace()
					element.Tag = err.Tag()
					element.Value = err.Param()
					errors = append(errors, &element)
			}
	}
	return errors
}
func connectDatabase() (*deta.Deta) {
	d, err := deta.New(deta.WithProjectKey("b074ax0y_QmjZ3p9MmsJVbYp3d7vkGeeWNjxHU6yW"))
	if err != nil {
			fmt.Println("failed to init new Deta instance:", err)
			return d
		}
	return d
}
func main() {
    app := fiber.New()
		app.Use(cors.New(cors.Config{
			AllowOrigins: "*",
		}))
    // app.Get("/", helloWorldHandler)
		app.Get("/users", getUsers)
		app.Get("/user/:key", getUser)
		app.Post("/user", AddUser)
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
func getUser(c *fiber.Ctx) error {
	d := connectDatabase()
	param := c.AllParams()
	var paramId = ""
	for paramName, paramValue := range param {
		if paramName == "key" {
			paramId = paramValue
		}
	}
	db, err := base.New(d, "base_name")
	var users []*User
	_, err = db.Fetch(&base.FetchInput{
		Q: base.Query{
			{"key": paramId},
		},
		Dest: &users,

	})
	if err != nil {
			fmt.Println("failed to fetch item:", err)
			return c.Status(fiber.StatusBadRequest).JSON("El usuario no existe")
	}
	fmt.Print(users)
	return c.JSON(users)
}
// func getUser(c *fiber.Ctx) error {
// 	fmt.Print(c.AllParams()) // "{"name": "fenny"}"
// 	return c.JSON(fiber.Map{
// 		"name": "John Doe",
// 	})
// 	// ...
// }

func getUsers(c *fiber.Ctx) error {
		d := connectDatabase()
		db, err := base.New(d, "base_name")
			var users []*User
	_, err = db.Fetch(&base.FetchInput{
			Dest: &users,
	})
	fmt.Println(users)
	if err != nil {
			fmt.Println("failed to fetch items:", err)
			return c.Status(fiber.StatusBadRequest).JSON("Error al obtener los usuarios")
	}

		return c.JSON(users)
}
func AddUser(c *fiber.Ctx) error {
	//Connect to database
	d := connectDatabase()

	user := new(User)

	if err := c.BodyParser(user); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"message": err.Error(),
			})
		 
	}

	errors := ValidateStruct(*user)
	if errors != nil {
		 return c.Status(fiber.StatusBadRequest).JSON(errors)
			
	}
	db, err := base.New(d, "base_name")
	_, err = db.Insert(&User{
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