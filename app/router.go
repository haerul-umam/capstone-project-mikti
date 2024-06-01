package app

import (
	"log"

	"github.com/haerul-umam/capstone-project-mikti/controller"
	"github.com/haerul-umam/capstone-project-mikti/helper"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	customMiddleware "github.com/haerul-umam/capstone-project-mikti/middleware"
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
	e.Use(middleware.Recover())
	e.Validator = helper.NewValidator()
	e.HTTPErrorHandler = helper.BindAndValidate

	// Auth Controller
	e.POST("/v1/login", authController.Login)
	e.POST("/v1/register", authController.Register)

	adminRoutes := e.Group("/api/admin")
	adminRoutes.Use(customMiddleware.JWTProtection())
	adminRoutes.Use(customMiddleware.JWTAuthRole("ADMIN"))

	buyerRoutes := e.Group("/api")
	buyerRoutes.Use(customMiddleware.JWTProtection())
	buyerRoutes.Use(customMiddleware.JWTAuthRole("BUYER"))

	return e
}
