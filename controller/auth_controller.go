package controller

import "github.com/labstack/echo/v4"

type AuthController interface {
	Login(e echo.Context) error
}