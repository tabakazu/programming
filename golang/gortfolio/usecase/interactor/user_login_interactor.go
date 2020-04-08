package interactor

import (
	"github.com/tabakazu/gortfolio/domain/entity"
	"github.com/tabakazu/gortfolio/domain/repository"
	"github.com/tabakazu/gortfolio/usecase"
	"github.com/tabakazu/gortfolio/usecase/requests"
	"golang.org/x/crypto/bcrypt"
)

type UserLoginInteractor struct {
	repository repository.UserRepository
}

func NewUserLoginInteractor(repo repository.UserRepository) usecase.UserLoginUseCase {
	return &UserLoginInteractor{
		repository: repo,
	}
}

func (u UserLoginInteractor) Call(req requests.UserLoginRequest) (entity.User, error) {
	user, err := u.repository.FindByEmail(req.Email)
	if err != nil {
		return user, err
	}
	if err := authenticate(req.Password, user.PasswordDigest); err != nil {
		return user, err
	}
	return user, nil
}

func authenticate(passwd, passwdDigest string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(passwdDigest), []byte(passwd)); err != nil {
		return err
	}
	return nil
}
