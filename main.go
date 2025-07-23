package main

import (
	"fmt"
	"log"

	"go-fiber-todo/user"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name  string
	Email string
}

func main() {
	dsn := "postgresql://neondb_owner:npg_agmowK25PsHZ@ep-long-bird-a1b82q7u-pooler.ap-southeast-1.aws.neon.tech/neondb?sslmode=require&channel_binding=require"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect dababase :", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("ping err", err)
	}

	err = sqlDB.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("success connect db")

	err = db.AutoMigrate(&User{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("✅ Migration succeeded.")

	app := fiber.New()
	user.RegisterUserRoutes(app, db)
	log.Fatal(app.Listen(":3000"))
}
