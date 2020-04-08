package usecase

import (
	"github.com/tabakazu/gortfolio/domain/entity"
	"github.com/tabakazu/gortfolio/usecase/requests"
)

type UserLoginUseCase interface {
	Call(req requests.UserLoginRequest) (entity.User, error)
}
