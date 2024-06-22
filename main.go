package main

import (
	"fmt"
	"os"

	"github.com/haerul-umam/capstone-project-mikti/app"
	"github.com/haerul-umam/capstone-project-mikti/helper"
)

func main() {
	helper.ValidateEnv()

	env := os.Getenv("ENV")
	fmt.Println(env, "ini")
	if env == "production" {
		app.Migrate()
	}

	server := StartServer()
	if err := server.Start(":8000"); err != nil {
		return
	}
}