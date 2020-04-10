package usecase

import (
	"github.com/tabakazu/gortfolio/usecase/requests"
)

type UserRegisterUseCase interface {
	Call(req requests.UserRegisterRequest) error
}
