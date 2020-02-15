package usecase

import (
	"github.com/tabakazu/practice20191209/domain/repository"
)

type UserUsecase struct {
	UserRepository repository.UserRepository
}
