package usecase_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/tabakazu/practice20200105/application/usecase"
	"github.com/tabakazu/practice20200105/domain/model"
	"github.com/tabakazu/practice20200105/mocks"
)

func TestGetItem(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	itemRepository := mocks.NewMockItemRepository(mockCtrl)
	u := &usecase.ItemUsecase{Repository: itemRepository}

	itemId := uint(1)
	record := &model.Item{Model: model.Model{ID: itemId}}
	itemRepository.EXPECT().FindByID(uint(1)).Return(record, nil)

	item, _ := u.GetItem(itemId)
	if item.ID != itemId {
		t.Errorf("GetItem(%v) = %v want %v", itemId, item.ID, itemId)
	}
}
