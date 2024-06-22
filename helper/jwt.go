package helper

import (
	"errors"

	"github.com/golang-jwt/jwt/v5"
	"github.com/haerul-umam/capstone-project-mikti/config"
	"github.com/labstack/echo/v4"
)

type JwtClaims struct {
	ID 		string   `json:"user_id"`
	Email string 	 `json:"email"`
	Role  string   `json:"role"`
	jwt.RegisteredClaims
}

type TokenUseCase interface {
	GenerateAccessToken(claims JwtClaims) (string, error)
}

type TokenUseCaseImpl struct{}

func NewTokenUseCase() *TokenUseCaseImpl {
	return &TokenUseCaseImpl{}
}

func (t *TokenUseCaseImpl) GenerateAccessToken(claims JwtClaims) (string, error) {
	plainToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	encodedToken, err := plainToken.SignedString([]byte(config.GetEnv().SecretKey))
	if err != nil {
		return "", err
	}

	return encodedToken, nil
}

func GetClaimsValue(c echo.Context) (JwtClaims) {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*JwtClaims)
	return *claims
} 

func ValidateRoleJWT(c echo.Context, role string) error {
	claims := GetClaimsValue(c)

	if claims.Role != role {
		return errors.New("forbidden resource")
	}
	return nil
}
