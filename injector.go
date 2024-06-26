//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/haerul-umam/capstone-project-mikti/app"
	"github.com/haerul-umam/capstone-project-mikti/controller"
	"github.com/haerul-umam/capstone-project-mikti/helper"
	"github.com/haerul-umam/capstone-project-mikti/repository"
	"github.com/haerul-umam/capstone-project-mikti/service"
	"github.com/labstack/echo/v4"
)

var authSet = wire.NewSet(
	repository.NewUserRepository,
	wire.Bind(new(repository.UserRepository), new(*repository.UserRepositoryImpl)),
	helper.NewTokenUseCase,
	wire.Bind(new(helper.TokenUseCase), new(*helper.TokenUseCaseImpl)),
	service.NewUserService,
	wire.Bind(new(service.UserService), new(*service.UserServiceImpl)),
	controller.NewAuthController,
	wire.Bind(new(controller.AuthController), new(*controller.AuthControllerImpl)),
)

var orderSet = wire.NewSet(
	repository.NewOrderRepository,
	wire.Bind(new(repository.OrderRepository), new(*repository.OrderRepositoryImpl)),
	service.NewOrderService,
	wire.Bind(new(service.OrderService), new(*service.OrderServiceImpl)),
	controller.NewOrderController,
	wire.Bind(new(controller.OrderController), new(*controller.OrderControllerImpl)),
)

var categorySet = wire.NewSet(
	repository.NewCategoryRepository,
	wire.Bind(new(repository.CategoryRepository), new(*repository.CategoryRepositoryImpl)),
	service.NewCategoryService,
	wire.Bind(new(service.CategoryService), new(*service.CategoryServiceImpl)),
	controller.NewCategoryController,
	wire.Bind(new(controller.CategoryEventController), new(*controller.CategoryContollerImpl)),
)

var eventSet = wire.NewSet(
	repository.NewEventRepository,
	wire.Bind(new(repository.EventRepository), new(*repository.EventRepositoryImpl)),
	service.NewEventService,
	wire.Bind(new(service.EventService), new(*service.EventServiceImpl)),
	controller.NewEventController,
	wire.Bind(new(controller.EventController), new(*controller.EventControllerImpl)),
)

func StartServer() *echo.Echo {
	wire.Build(
		app.InitConnetion,
		authSet,
		orderSet,
		categorySet,
		eventSet,
		app.Router,
	)
	return nil
}
