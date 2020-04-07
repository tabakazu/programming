package interactor

import (
	"github.com/tabakazu/gortfolio/domain/entity"
	"github.com/tabakazu/gortfolio/domain/repository"
	"github.com/tabakazu/gortfolio/usecase"
	"github.com/tabakazu/gortfolio/usecase/requests"
	"golang.org/x/crypto/bcrypt"
)

type UserRegisterInteractor struct {
	repository repository.UserRepository
}

func NewUserRegisterInteractor(repo repository.UserRepository) usecase.UserCreateUseCase {
	return &UserRegisterInteractor{
		repository: repo,
	}
}

func (u UserRegisterInteractor) Call(req requests.UserCreateRequest) error {
	user := &entity.User{
		Email:    req.Email,
		Password: req.Password,
	}

	if err := encryptPassword(user); err != nil {
		return err
	}

	if err := u.repository.Create(user); err != nil {
		return err
	}

	return nil
}

func encryptPassword(user *entity.User) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.PasswordDigest = string(hash)
	return nil
}
