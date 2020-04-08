package requests

type UserLoginRequest struct {
	Email    string `json: "email"`
	Password string `json: "password"`
}

func NewUserLoginRequest() UserLoginRequest {
	return UserLoginRequest{}
}
