package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDb(das string) *gorm.DB {
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
