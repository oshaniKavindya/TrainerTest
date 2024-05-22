package main

import (
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// User represents a user in the system
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

// Global variable to store user data (in-memory storage for demonstration purposes)
var users []User

func main() {
	// Initialize Fiber app
	app := fiber.New()

	// Define routes
	app.Post("/users", createUser)
	app.Get("/users/:id", getUser)
	app.Put("/users/:id", updateUser)
	app.Delete("/users/:id", deleteUser)

	// Start server
	log.Fatal(app.Listen(":3000"))
}

// Handler to create a new user
func createUser(c *fiber.Ctx) error {
	// Parse request body
	var newUser User
	if err := c.BodyParser(&newUser); err != nil {
		return err
	}

	// Assign ID to the new user
	newUser.ID = len(users) + 1

	// Add user to the slice
	users = append(users, newUser)

	// Return success response
	return c.Status(fiber.StatusCreated).JSON(newUser)
}

// Handler to get a user by ID
func getUser(c *fiber.Ctx) error {
	// Extract user ID from route parameters
	userID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}

	// Find user by ID
	var foundUser *User
	for _, user := range users {
		if user.ID == userID {
			foundUser = &user
			break
		}
	}

	// Check if user is found
	if foundUser == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "User not found"})
	}

	// Return user data
	return c.Status(fiber.StatusOK).JSON(foundUser)
}

// Handler to update a user by ID
func updateUser(c *fiber.Ctx) error {
	// Extract user ID from route parameters
	userID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}

	// Find user index by ID
	var userIndex int = -1
	for i, user := range users {
		if user.ID == userID {
			userIndex = i
			break
		}
	}

	// Check if user exists
	if userIndex == -1 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "User not found"})
	}

	// Parse request body
	var updatedUser User
	if err := c.BodyParser(&updatedUser); err != nil {
		return err
	}

	// Update user data
	users[userIndex] = updatedUser

	// Return success response
	return c.Status(fiber.StatusOK).JSON(updatedUser)
}

// Handler to delete a user by ID
func deleteUser(c *fiber.Ctx) error {
	// Extract user ID from route parameters
	userID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}

	// Find user index by ID
	var userIndex int = -1
	for i, user := range users {
		if user.ID == userID {
			userIndex = i
			break
		}
	}

	// Check if user exists
	if userIndex == -1 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "User not found"})
	}

	// Remove user from the slice
	users = append(users[:userIndex], users[userIndex+1:]...)

	// Return success response
	c.SendStatus(fiber.StatusNoContent)
	return nil
}
