package routes

import (
	"github.com/haerul-umam/capstone-project-mikti/controller"
	"github.com/labstack/echo/v4"
)

func PublicRoutes(
	e *echo.Echo,
	authController controller.AuthController,
	eventController controller.EventController,
) {
	public := e.Group("/api")

	public.POST("/v1/login", authController.Login)
	public.POST("/v1/register", authController.Register)

	public.GET("/v1/event", eventController.GetAllEvents)
	public.GET("/v1/event/:event_id", eventController.GetEvent)
}