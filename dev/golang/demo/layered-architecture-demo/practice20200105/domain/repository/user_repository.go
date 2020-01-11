package repository

import "github.com/tabakazu/practice20200105/domain/model"

type UserRepository interface {
	FindByEmail(email string) (*model.User, error)
}
