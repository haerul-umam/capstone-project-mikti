package helper

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/haerul-umam/capstone-project-mikti/model/web"
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
				report.Message = fmt.Sprintf("%s tidak valid", err.Field())
				report.Code = 400
			case "min":
				report.Message = fmt.Sprintf("%s harus terdiri dari minimal %s karakter", err.Field(), err.Param())
				report.Code = 400
			case "status_check":
				report.Message = fmt.Sprintf("%s pembayaran tidak valid", err.Field())
				report.Code = 400
			case "gte":
				report.Message = fmt.Sprintf("%s harus lebih atau sama dengan %s", err.Field(), err.Param())
				report.Code = 400
			case "gt":
				report.Message = fmt.Sprintf("%s harus lebih dari %s", err.Field(), err.Param())
				report.Code = 400
			default:
				report.Message = fmt.Sprintf("%s field tidak valid", err.Field())
				report.Code = 400
			}
		}
	}

	c.Logger().Error(report.Message)
	c.JSON(report.Code, web.ResponseToClient(report.Code, report.Message.(string), nil))
}
