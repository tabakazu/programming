package usecase

import (
	"github.com/tabakazu/simple-auth-server-demo/domain/entry"
	"github.com/tabakazu/simple-auth-server-demo/infra"
	"github.com/tabakazu/simple-auth-server-demo/infra/repo"
)

type UserUsecase struct{}

func NewUserUsecase() *UserUsecase {
	return &UserUsecase{}
}

func (u *UserUsecase) Register(email, passwd string) (*entry.User, error) {
	passwdDigest, _ := infra.GeneratePasswordDigest(passwd)
	apiToken, err := infra.GenerateAPIToken()

	r := repo.NewUserRepo()
	user, err := r.Create(email, passwdDigest, apiToken)

	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserUsecase) Authenticate(email, passwd string) (*entry.User, error) {
	r := repo.NewUserRepo()
	user, err := r.FindByEmail(email)
	if err != nil {
		return nil, err
	}

	if err := infra.ComparePassword(passwd, user.PasswordDigest); err != nil {
		return nil, err
	}

	return user, nil
}
