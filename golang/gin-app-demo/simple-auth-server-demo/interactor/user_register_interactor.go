package interactor

import (
	"github.com/tabakazu/simple-auth-server-demo/domain"
	"github.com/tabakazu/simple-auth-server-demo/domain/repository"
	"github.com/tabakazu/simple-auth-server-demo/infra"
	"github.com/tabakazu/simple-auth-server-demo/infra/db"
)

type UserRegisterInteractor struct {
	Repo repository.UserRepository
}

func NewUserRegisterInteractor() *UserRegisterInteractor {
	return &UserRegisterInteractor{
		Repo: db.NewUserRepository(),
	}
}

func (i *UserRegisterInteractor) Handle(email, passwd string) (*domain.User, error) {
	passwdDigest, _ := infra.GeneratePasswordDigest(passwd)
	apiToken, err := infra.GenerateAPIToken()

	user, err := i.Repo.Create(email, passwdDigest, apiToken)

	if err != nil {
		return nil, err
	}
	return user, nil
}
