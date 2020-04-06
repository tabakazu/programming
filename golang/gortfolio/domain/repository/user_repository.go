package repository

import "github.com/tabakazu/gortfolio/domain/entity"

type UserRepository interface {
	Create(user *entity.User) error
}
