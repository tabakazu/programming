package usecase_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/tabakazu/gomock-demo/domain"
	"github.com/tabakazu/gomock-demo/mocks"
	"github.com/tabakazu/gomock-demo/usecase"
)

func TestGetPostByID(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepository := mocks.NewMockPostRepository(mockCtrl)
	testUsecase := usecase.PostUsecase{Repository: mockRepository}

	mockRepository.EXPECT().FindByID(1).Return(&domain.Post{}, nil).Times(1)
	testUsecase.GetPostByID(1)
}
