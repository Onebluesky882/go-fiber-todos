package main

import (
	"fmt"
	"log"

	"go-fiber-todo/database"
	"go-fiber-todo/user"

	"github.com/gofiber/fiber/v2"
)

func main() {
	db := database.ConnectDb()

	err := db.AutoMigrate(&user.User{})
	if err != nil {
		log.Fatalf("AutoMigrate error: %v", err)
	}
	fmt.Println("✅ Migration succeeded.")

	app := fiber.New()
	user.RegisterUserRoutes(app, db)
	log.Fatal(app.Listen(":3000"))
}
