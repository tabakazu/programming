package usecase_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/tabakazu/interface-pattern-demo/domain"
	"github.com/tabakazu/interface-pattern-demo/mocks"
	"github.com/tabakazu/interface-pattern-demo/usecase"
)

func Test_GetPosts(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mock := mocks.NewMockPostRepository(mockCtrl)
	usecase := usecase.PostUsecase{PostRepository: mock}

	mock.EXPECT().All().Return([]domain.Post{}, nil).Times(1)
	usecase.GetPosts()
}

func Test_GetPost(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mock := mocks.NewMockPostRepository(mockCtrl)
	usecase := usecase.PostUsecase{PostRepository: mock}

	mock.EXPECT().FindById(1).Return(domain.Post{}, nil).Times(1)
	usecase.GetPost(1)
}

func Test_CreatePost(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mock := mocks.NewMockPostRepository(mockCtrl)
	usecase := usecase.PostUsecase{PostRepository: mock}

	post := domain.Post{
		Title: "Test Title",
		Body:  "Test Body",
	}

	mock.EXPECT().Store(post).Return(domain.Post{}, nil).Times(1)
	usecase.CreatePost(post)
}

func Test_UpdatePost(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mock := mocks.NewMockPostRepository(mockCtrl)
	usecase := usecase.PostUsecase{PostRepository: mock}

	post := domain.Post{
		Title: "Test Title",
		Body:  "Test Body",
	}

	mock.EXPECT().Update(1, post).Return(domain.Post{}, nil).Times(1)
	usecase.UpdatePost(1, post)
}

func Test_DestroyPost(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mock := mocks.NewMockPostRepository(mockCtrl)
	usecase := usecase.PostUsecase{PostRepository: mock}

	mock.EXPECT().Destroy(1).Return(nil).Times(1)
	usecase.DestroyPost(1)
}
