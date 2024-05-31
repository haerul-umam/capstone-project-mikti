package controller

import "github.com/labstack/echo/v4"

type OrderController interface {
	CreateOrder(e echo.Context) error
}
