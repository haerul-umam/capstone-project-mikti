package middleware

import (
	"net/http"
	"os"

	"github.com/haerul-umam/tugas-mikti-kel1/helper"
	"github.com/haerul-umam/tugas-mikti-kel1/model/web"
	"github.com/labstack/echo/v4"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
)

func JWTProtection() echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(helper.JwtClaims)
		},
		SigningKey: []byte(os.Getenv("SECRET_KEY")),
		ErrorHandler: func(c echo.Context, err error) error {
			return c.JSON(
				http.StatusUnauthorized,
				web.ResponseToClient(http.StatusUnauthorized, "akses ditolak", nil),
			)
		},
	})
}