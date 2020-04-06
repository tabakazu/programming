package interactor

import (
	"github.com/tabakazu/gortfolio/domain/entity"
	"github.com/tabakazu/gortfolio/domain/repository"
	"github.com/tabakazu/gortfolio/usecase"
	"github.com/tabakazu/gortfolio/usecase/requests"
)

type UserCreateInteractor struct {
	repository repository.UserRepository
}

func NewUsersInteractor(repo repository.UserRepository) usecase.UserCreateUseCase {
	return &UserCreateInteractor{
		repository: repo,
	}
}

func (u UserCreateInteractor) Call(req requests.UserCreateRequest) error {
	user := &entity.User{
		Email: req.Email,
	}
	if err := u.repository.Create(user); err != nil {
		return err
	}
	return nil
}
