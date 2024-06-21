package app

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitConnetion() *gorm.DB {
	dsn := os.Getenv("DSN")
	configuration := gorm.Config{}

	db, err := gorm.Open(postgres.Open(dsn), &configuration)
	if err != nil {
		log.Fatal(err.Error())
	}

	return db
}