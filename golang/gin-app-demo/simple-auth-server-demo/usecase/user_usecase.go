package usecase

import (
	"github.com/tabakazu/simple-auth-server-demo/domain/entry"
)

type UserUsecase struct{}

func NewUserUsecase() *UserUsecase {
	return &UserUsecase{}
}

func (u *UserUsecase) Register(email, passwd string) *entry.User {
	return &entry.User{
		Email:             email,
		EncryptedPassword: passwd,
	}
}
