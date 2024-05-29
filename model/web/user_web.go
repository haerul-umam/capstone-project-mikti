package web

type UserLoginRequest struct {
	Password string `validate:"required" json:"password"`
	Email    string `validate:"email,required" json:"email"`
}

type UserLoginResponse struct {
	Token string `json:"token"`
	Email string `json:"email"`
	Role 	string `json:"role"`
}

type UserRegisterRequest struct {
	Name     string `validate:"required" json:"name"`
	Email    string `validate:"required,email" json:"email"`
	Password string `validate:"required" json:"password"`
}

type UserRegisterResponse struct {
	Token string `json:"token"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
