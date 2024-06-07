package controller

import "github.com/labstack/echo/v4"

type CategoryEventController interface {
	NewCategory(e echo.Context) error
	GetCategoryList(c echo.Context) error
	UpdateCategory(c echo.Context) error
	DeleteCategory(c echo.Context) error
}
