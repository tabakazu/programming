package interactor

import (
	"github.com/tabakazu/simple-auth-server-demo/domain"
	"github.com/tabakazu/simple-auth-server-demo/domain/repository"
	"github.com/tabakazu/simple-auth-server-demo/infra"
	"github.com/tabakazu/simple-auth-server-demo/infra/db"
)

type UserAuthenticateInteractor struct {
	Repo repository.UserRepository
}

func NewUserAuthenticateInteractor() *UserAuthenticateInteractor {
	return &UserAuthenticateInteractor{
		Repo: db.NewUserRepository(),
	}
}

func (i *UserAuthenticateInteractor) Handle(email, passwd string) (*domain.User, error) {
	user, err := i.Repo.FindByEmail(email)
	if err != nil {
		return nil, err
	}

	if err := infra.ComparePassword(passwd, user.PasswordDigest); err != nil {
		return nil, err
	}

	return user, nil
}
