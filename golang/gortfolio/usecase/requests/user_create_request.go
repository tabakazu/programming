package requests

type UserCreateRequest struct {
	Email    string `json: "email"`
	Password string `json: "password"`
}

func NewUserCreateRequest() UserCreateRequest {
	return UserCreateRequest{}
}
