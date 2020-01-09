package repository

import "github.com/tabakazu/practice20200105/domain/model"

type ItemRepository interface {
	FindByID(id uint) (*model.Item, error)
}
