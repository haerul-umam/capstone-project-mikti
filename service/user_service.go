package service

import "github.com/haerul-umam/capstone-project-mikti/model/web"

type UserService interface {
	LoginUser(email, password string) (web.UserLoginResponse, error)
	SaveUser(request web.UserRegisterRequest) (web.UserRegisterResponse, error)
}
