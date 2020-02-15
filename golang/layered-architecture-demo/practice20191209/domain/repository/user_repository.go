package repository

import (
	"github.com/tabakazu/practice20191209/domain/model"
)

type UserRepository interface {
	Save() (model.User, error)
}
