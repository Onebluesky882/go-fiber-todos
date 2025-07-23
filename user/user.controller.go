package user

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	Service *UserService
}

// head
func NewUserController(service *UserService) *UserController {
	return &UserController{Service: service}
}

// post / new user
func (ctl *UserController) CreateUser(c *fiber.Ctx) error {
	var input User
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	if err := ctl.Service.Create(&input); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "cannot create user"})
	}
	fmt.Printf("Created user: %+v\n", input)
	return c.Status(fiber.StatusCreated).JSON(input)
}
