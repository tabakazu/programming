package usecase

import "github.com/tabakazu/simple-auth-server-demo/domain"

type UserRegisterUseCase interface {
	Handle(email, passwd string) (*domain.User, error)
}

type UserAuthenticateUseCase interface {
	Handle(email, passwd string) (*domain.User, error)
}

type UserGetByTokenUseCase interface {
	Handle(token string) (*domain.User, error)
}
