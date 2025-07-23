package user

import (
	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	Service *UserService
}

func NewUserController(s *UserService) *UserController {
	return &UserController{Service: s}
}

func (ctl *UserController) GetUsers(c *fiber.Ctx) error {
	users, err := ctl.Service.FindAll()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "cannot fetch users"})
	}
	return c.JSON(users)
}

func (ctl *UserController) CreateUser(c *fiber.Ctx) error {
	var user User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid request"})
	}
	if err := ctl.Service.Create(&user); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "cannot create user"})
	}
	return c.Status(201).JSON(user)
}
