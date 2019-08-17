package controller_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/tabakazu/gomock-demo/controller"
	"github.com/tabakazu/gomock-demo/mocks"
)

func TestUse(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockUsecase := mocks.NewMockPost(mockCtrl)
	testController := &controller.Post{Usecase: mockUsecase}

	mockUsecase.EXPECT().FindByID(1).Return(nil).Times(1)
	testController.Show()
}
