package helper

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/haerul-umam/tugas-mikti-kel1/model/web"
	"github.com/labstack/echo/v4"
)

func BindAndValidate(err error, c echo.Context) {
	report, ok := err.(*echo.HTTPError)

	if !ok {
		report = echo.NewHTTPError(500, err.Error())
	}

	if castedObject, ok := err.(validator.ValidationErrors); ok {
		for _, err := range castedObject {
			switch err.Tag() {
			case "required":
				report.Message = fmt.Sprintf("%s field tidak boleh kosong", err.Field())
				report.Code = 400
			case "email":
				report.Message = fmt.Sprintf("%s email tidak valid", err.Field())
				report.Code = 40
			}
		}
	}

	c.Logger().Error(report.Message)
	c.JSON(report.Code, report.Message.(string))
	c.JSON(report.Code, web.ResponseToClient(report.Code, report.Message.(string), nil))
}