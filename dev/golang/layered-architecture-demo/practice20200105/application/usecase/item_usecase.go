package usecase

import (
	"github.com/tabakazu/practice20200105/domain/model"
	"github.com/tabakazu/practice20200105/domain/repository"
)

type ItemUsecase struct {
	Repository repository.ItemRepository
}

func (u *ItemUsecase) GetItem(id uint) (*model.Item, error) {
	item, err := u.Repository.FindByID(id)
	if err != nil {
		return nil, err
	}
	return item, nil
}
