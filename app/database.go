package app

import (
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/haerul-umam/capstone-project-mikti/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


func InitConnetion() *gorm.DB {
	var dsn = config.GetEnv().DSN
	configuration := gorm.Config{}

	db, err := gorm.Open(postgres.Open(dsn), &configuration)
	if err != nil {
		log.Fatal(err.Error())
	}

	return db
}

func Migrate() {
	var dsn = config.GetEnv().DSN
	m, err := migrate.New("file://migrations", dsn)

	if err != nil {
		log.Fatal(err)
	}

	if err := m.Up(); err != nil {
		log.Println(err)
	}
}