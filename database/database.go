package database

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDb() *gorm.DB {
	// log .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Warning : .env file not found , using system env variables")
	}
	// find DB="adsgfahg"
	das := os.Getenv("DB")
	if das == "" {
		log.Fatal(("DB environment variable is not set"))
	}

	db, err := gorm.Open(postgres.Open(das), &gorm.Config{})
	if err != nil {
		log.Fatal("fail to connect databae ")
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("failed to ping db: " + err.Error())
	}

	err = sqlDB.Ping()
	if err != nil {
		log.Fatal("failed to ping db: " + err.Error())
	}
	return db
}
