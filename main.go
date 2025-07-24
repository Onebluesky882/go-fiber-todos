package main

import (
	"log"

	"go-fiber-todo/config"
	"go-fiber-todo/database"
	"go-fiber-todo/posts"
	"go-fiber-todo/user"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config.LoadEnv()
	dsn := config.GetEnv("DB")
	db := database.ConnectDb(dsn)

	err := db.AutoMigrate(&user.User{})
	if err != nil {
		log.Fatalf("AutoMigrate error: %v", err)
	}

	app := fiber.New()

	posts.RegisterPostRouted(app)
	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("pong")
	})
	user.RegisterUserRoutes(app, db)
	log.Fatal(app.Listen(":3000"))
}
