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