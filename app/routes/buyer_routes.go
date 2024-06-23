package routes

import (
	"github.com/haerul-umam/capstone-project-mikti/controller"
	"github.com/haerul-umam/capstone-project-mikti/middleware"
	"github.com/labstack/echo/v4"
)

func BuyerRoutes(
	e *echo.Echo,
	orderController controller.OrderController,
) {
	buyer := e.Group("/api")
	buyer.Use(middleware.JWTProtection())
	buyer.Use(middleware.JWTAuthRole("BUYER"))

	buyer.POST("/v1/order", orderController.CreateOrder)
	buyer.GET("/v1/detail/:id", orderController.DetailOrder)
}