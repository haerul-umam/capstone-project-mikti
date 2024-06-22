package midleware

import (
	"net/http"

	"github.com/haerul-umam/capstone-project-mikti/config"
	"github.com/haerul-umam/capstone-project-mikti/helper"
	"github.com/haerul-umam/capstone-project-mikti/model/web"
	"github.com/labstack/echo/v4"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
)

func JWTProtection() echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(helper.JwtClaims)
		},
		SigningKey: []byte(config.GetEnv().SecretKey),
		ErrorHandler: func(c echo.Context, err error) error {
			return c.JSON(
				http.StatusUnauthorized,
				web.ResponseToClient(http.StatusUnauthorized, "akses ditolak", nil),
			)
		},
	})
}

func JWTAuthRole(role string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
	
			err := helper.ValidateRoleJWT(c, role)
			if err != nil {
				return c.JSON(
					http.StatusUnauthorized,
					web.ResponseToClient(http.StatusUnauthorized, err.Error(), nil),
				)
			}
	
			return next(c)
		}
	}
}