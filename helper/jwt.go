package helper

import (
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type JwtClaims struct {
	ID    string `json:"user_id"`
	Email string `json:"email"`
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

	encodedToken, err := plainToken.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return "", err
	}

	return encodedToken, nil
}

func GetClaimsValue(c echo.Context) JwtClaims {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*JwtClaims)
	return *claims
}
