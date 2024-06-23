package routes

import (
	"github.com/haerul-umam/capstone-project-mikti/controller"
	"github.com/labstack/echo/v4"
	"github.com/haerul-umam/capstone-project-mikti/middleware"
)

func AdminRoutes(
	e *echo.Echo,
	eventController controller.EventController,
	orderController controller.OrderController,
	categoryContoller controller.CategoryEventController,
) {
	admin := e.Group("/api/admin")
	admin.Use(middleware.JWTProtection())
	admin.Use(middleware.JWTAuthRole("ADMIN"))

	admin.GET("/v1/order", orderController.GetOrdersPage)

	admin.PATCH("/v1/event/:event_id", eventController.UpdateEvent)
	admin.DELETE("/v1/event/:event_id", eventController.DeleteEvent)
	admin.GET("/v1/event/:event_id", eventController.GetEventAdmin)
	admin.POST("/v1/event", eventController.CreateEvent)

	admin.POST("/v1/payment/:id/status", orderController.ChangeOrderStatus)
	admin.GET("/v1/payment", orderController.GetAllPayment)

	admin.POST("/v1/category", categoryContoller.NewCategory)
	admin.GET("/v1/categories", categoryContoller.GetCategoryList)
	admin.PATCH("/v1/category/:id", categoryContoller.UpdateCategory)
	admin.DELETE("/v1/category/:id", categoryContoller.DeleteCategory)
}