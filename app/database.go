package app

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"github.com/joho/godotenv"
)

func InitConnetion() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading env file")
	}

	dsn := os.Getenv("DSN")
	configuration := gorm.Config{}

	db, err := gorm.Open(postgres.Open(dsn), &configuration)
	if err != nil {
		log.Fatal(err.Error())
	}

	return db
}