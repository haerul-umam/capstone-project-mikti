package app

import (
	"log"

	"github.com/haerul-umam/capstone-project-mikti/controller"
	"github.com/haerul-umam/capstone-project-mikti/helper"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Router(
	authController controller.AuthController,
) *echo.Echo {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading env file")
	}

	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	e.Validator = helper.NewValidator()
	e.HTTPErrorHandler = helper.BindAndValidate

	// Auth Controller
	e.POST("/v1/login", authController.Login)

	return e
}