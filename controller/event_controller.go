package controller

import "github.com/labstack/echo/v4"

type EventController interface {
	GetEvent(c echo.Context) error
	UpdateEvent(c echo.Context) error
}
