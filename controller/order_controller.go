package controller

import "github.com/labstack/echo/v4"

type OrderController interface {
	CreateOrder(e echo.Context) error
	GetOrdersPage(e echo.Context) error
	DetailOrder(e echo.Context) error
}
