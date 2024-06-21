package main

import "github.com/haerul-umam/capstone-project-mikti/helper"

func main() {
	helper.ValidateEnv()

	server := StartServer()
	if err := server.Start(":8000"); err != nil {
		return
	}
}