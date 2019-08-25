package usecase

import (
	"github.com/tabakazu/goapp-practice/domain/model"
	"github.com/tabakazu/goapp-practice/domain/repository"
	"github.com/tabakazu/goapp-practice/infrastructure/datastore"
)

type PostUsecase struct {
	Repository repository.PostRepository
}

func NewPostUsecase() *PostUsecase {
	return &PostUsecase{
		Repository: datastore.NewPostRepository(),
	}
}

func (u *PostUsecase) GetPostByID(id int) (*model.Post, error) {
	post, err := u.Repository.FindByID(id)
	if err != nil {
		return nil, err
	}
	return post, nil
}

func (u *PostUsecase) CreatePost(p model.Post) (*model.Post, error) {
	post, err := u.Repository.Create(p)
	if err != nil {
		return nil, err
	}
	return post, nil
}
