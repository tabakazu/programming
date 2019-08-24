package usecase_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/tabakazu/goapp-practice/application/usecase"
	"github.com/tabakazu/goapp-practice/domain/model"
	"github.com/tabakazu/goapp-practice/mocks"
)

func Test_PostUsecase_GetPost(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mock := mocks.NewMockPostRepository(mockCtrl)
	u := usecase.PostUsecase{Repository: mock}

	postID := 1

	mock.EXPECT().FindByID(postID).Return(&model.Post{}, nil).Times(1)
	u.GetPostByID(postID)
}

func Test_PostUsecase_CreatePost(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mock := mocks.NewMockPostRepository(mockCtrl)
	u := usecase.PostUsecase{Repository: mock}

	post := model.Post{
		Title: "Test Title",
		Body:  "Test Body",
	}

	mock.EXPECT().Create(post).Return(&model.Post{}, nil).Times(1)
	u.CreatePost(post)
}
