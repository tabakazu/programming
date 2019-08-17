package controller_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/tabakazu/gomock-demo/controller"
	"github.com/tabakazu/gomock-demo/domain"
	"github.com/tabakazu/gomock-demo/mocks"
)

func TestShow(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockUsecase := mocks.NewMockPostUsecase(mockCtrl)
	testController := controller.PostController{Usecase: mockUsecase}

	mockUsecase.EXPECT().GetPostByID(1).Return(&domain.Post{}, nil).Times(1)
	testController.Show()
}
