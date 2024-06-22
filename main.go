package main

import (
	"github.com/haerul-umam/capstone-project-mikti/app"
	"github.com/haerul-umam/capstone-project-mikti/config"
)

func main() {
	config.ValidateEnv()
	env := config.GetEnv().Env

	if env == "production" {
		app.Migrate()
	}

	server := StartServer()
	if err := server.Start(":8000"); err != nil {
		return
	}
}