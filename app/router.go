package app

import (
	"github.com/haerul-umam/capstone-project-mikti/controller"
	"github.com/haerul-umam/capstone-project-mikti/helper"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/haerul-umam/capstone-project-mikti/app/routes"
)

func Router(
	authController controller.AuthController,
	orderController controller.OrderController,
	eventController controller.EventController,

	categoryContoller controller.CategoryEventController,

) *echo.Echo {
	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Recover())
	e.Validator = helper.NewValidator()
	e.HTTPErrorHandler = helper.BindAndValidate

	routes.PublicRoutes(e, authController, eventController)
	routes.AdminRoutes(e, eventController, orderController, categoryContoller)
	routes.BuyerRoutes(e, orderController)

	return e
}
