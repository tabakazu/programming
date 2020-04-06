package usecase

import (
	"github.com/tabakazu/gortfolio/usecase/requests"
)

type UserCreateUseCase interface {
	Call(req requests.UserCreateRequest) error
}
