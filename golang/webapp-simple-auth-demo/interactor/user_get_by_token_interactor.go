package interactor

import (
	"github.com/tabakazu/simple-auth-server-demo/domain"
	"github.com/tabakazu/simple-auth-server-demo/domain/repository"
	"github.com/tabakazu/simple-auth-server-demo/infra/db"
)

type UserGetByTokenInteractor struct {
	Repo repository.UserRepository
}

func NewUserGetByTokenInteractor() *UserGetByTokenInteractor {
	return &UserGetByTokenInteractor{
		Repo: db.NewUserRepository(),
	}
}

func (i *UserGetByTokenInteractor) Handle(token string) (*domain.User, error) {
	user, err := i.Repo.FindByToken(token)
	if err != nil {
		return nil, err
	}
	return user, nil
}
