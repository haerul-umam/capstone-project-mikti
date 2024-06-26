package controller

import "github.com/labstack/echo/v4"

type EventController interface {
	CreateEvent(c echo.Context) error
	GetEvent(c echo.Context) error
	UpdateEvent(c echo.Context) error
	DeleteEvent(c echo.Context) error
	GetEventAdmin(c echo.Context) error
	GetAllEvents(e echo.Context) error
}
