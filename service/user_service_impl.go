package service

import (
	"errors"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/haerul-umam/capstone-project-mikti/helper"
	"github.com/haerul-umam/capstone-project-mikti/model/domain"
	"github.com/haerul-umam/capstone-project-mikti/model/web"
	"github.com/haerul-umam/capstone-project-mikti/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserServiceImpl struct {
	repo         repository.UserRepository
	tokenUseCase helper.TokenUseCase
}

func NewUserService(repo repository.UserRepository, token helper.TokenUseCase) *UserServiceImpl {
	return &UserServiceImpl{repo, token}
}

func (service *UserServiceImpl) LoginUser(email, password string) (web.UserLoginResponse, error) {
	user, err := service.repo.GetUserByEmail(email)

	if err != nil {
		return web.UserLoginResponse{}, errors.New("user atau password salah")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		return web.UserLoginResponse{}, errors.New("user atau password salah")
	}

	expired_minutes, _ := strconv.Atoi(os.Getenv("EXPIRED_TOKEN"))

	expired := time.Now().Local().Add(time.Duration(expired_minutes) * time.Minute)

	claims := &helper.JwtClaims{
		ID:    string(user.UserID),
		Email: user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "ticketing",
			ExpiresAt: jwt.NewNumericDate(expired),
		},
	}

	token, err := service.tokenUseCase.GenerateAccessToken(*claims)

	if err != nil {
		return web.UserLoginResponse{}, errors.New("server error: generating token")
	}

	return web.UserLoginResponse{
		Token: token,
		Email: user.Email,
		Role:  string(user.Role),
	}, nil
}

func (service *UserServiceImpl) SaveUser(request web.UserRegisterRequest) (web.UserRegisterResponse, error) {
	passHash, errHash := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.MinCost)
	if errHash != nil {
		return web.UserRegisterResponse{}, errHash
	}

	userRequest := domain.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: string(passHash),
		Role:     domain.Buyer,
	}

	saveUser, errSaveUser := service.repo.SaveUser(userRequest)

	if errSaveUser != nil {
		return web.UserRegisterResponse{}, errSaveUser
	}

	expired_minutes, _ := strconv.Atoi(os.Getenv("EXPIRED_TOKEN"))

	expired := time.Now().Local().Add(time.Duration(expired_minutes) * time.Minute)

	claims := &helper.JwtClaims{
		ID:    string(userRequest.UserID),
		Email: userRequest.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "ticketing",
			ExpiresAt: jwt.NewNumericDate(expired),
		},
	}

	token, err := service.tokenUseCase.GenerateAccessToken(*claims)

	if err != nil {
		return web.UserRegisterResponse{}, errors.New("server error: generating token")
	}

	return web.UserRegisterResponse{Token: token, Name: saveUser.Name, Email: saveUser.Email}, nil

}
