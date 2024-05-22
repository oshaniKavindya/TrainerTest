package main

import (
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)


type User struct {
    Id       int    `json:"id"`
    Name     string `json:"name"`
    Email    string `json:"email"`
    Age      int    `json:"age"`
}

var users = []User{
	{Id: 1, Name: "Anne", Email: "123@gmail.com", Age:20},
	{Id: 2, Name: "Josaph", Email: "josap@gmail.com" , Age :15},
	{Id: 3, Name: "Maria", Email: "M@gmail.com", Age:25},
}

func main() {
	app := fiber.New()


	app.Get("/users", func(c *fiber.Ctx) error {
		return c.Status(http.StatusOK).JSON(users)
	})

	app.Post("/users", func(c *fiber.Ctx) error {
		var newUser User

		if err := c.BodyParser(&newUser); err != nil {
			c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
			return err
		}
		users = append(users, newUser)
		return c.Status(http.StatusCreated).JSON(newUser)
	})

	app.Delete("/users/:id", func(c *fiber.Ctx) error {
		idStr := c.Params("id")
		id, _:= strconv.Atoi(idStr)

		for i, lang := range users {
			if lang.Id == id {
				users = append(users[:i], users[i+1:]...)
				break
			}
		}
		return c.Status(http.StatusNoContent).JSON(fiber.Map{"data": users})

	})

	app.Put("/users/:id", func(c *fiber.Ctx) error {
		idStr := c.Params("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
		}
		var updatedlanguage User
		updatedlanguage.Id = id 

		if err := c.BodyParser(&updatedlanguage); err != nil {
			c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
			return err
		}
		for i, lang := range users {
			if lang.Id == updatedlanguage.Id {
				users = append(users[:i], users[i+1:]...)
				users = append(users, updatedlanguage)
			}
		}
		return c.Status(http.StatusCreated).JSON(updatedlanguage)
	})

	app.Listen(":8080")
}
