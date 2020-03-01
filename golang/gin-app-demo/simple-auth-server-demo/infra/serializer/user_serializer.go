package serializer

import (
	"time"

	"github.com/tabakazu/simple-auth-server-demo/domain"
)

type UserSerializer struct {
	Email     string    `json:"email"`
	APIToken  string    `json:"api_token"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewUserSerializer(user *domain.User) *UserSerializer {
	return &UserSerializer{
		Email:    user.Email,
		APIToken: user.APIToken,
	}
}

func SerializeUser(user *domain.User) *UserSerializer {
	return NewUserSerializer(user)
}
