package user

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterUserRoutes(app *fiber.App, db *gorm.DB) {
	userService := NewUserService(db)
	userController := NewUserController(userService)

	route := app.Group("/users")
	route.Get("/", userController.GetUsers)
	route.Post("/", userController.CreateUser)
}
