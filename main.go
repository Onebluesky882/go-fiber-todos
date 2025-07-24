package main

import (
	"fmt"
	"log"

	"go-fiber-todo/config"
	"go-fiber-todo/database"
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
	fmt.Println("✅ Migration succeeded.")

	app := fiber.New()
	app.Get("/ping", func(c *fiber.Ctx) error {
		fmt.Println("Ping route called")
		return c.SendString("pong")
	})
	user.RegisterUserRoutes(app, db)
	log.Fatal(app.Listen(":3000"))
}
