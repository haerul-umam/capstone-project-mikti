package app

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
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

func Migrate() {
	dsn := os.Getenv("DSN")
	m, err := migrate.New("file://migrations", dsn)

	if err != nil {
		log.Fatal(err)
	}

	if err := m.Up(); err != nil {
		log.Println(err)
	}
}