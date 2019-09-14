package usecase

import "github.com/tabakazu/gomock-demo/domain"

type PostRepository interface {
	FindByID(id int) (*domain.Post, error)
}

type PostUsecase struct {
	Repository PostRepository
}

func (u PostUsecase) GetPostByID(id int) (*domain.Post, error) {
	post, err := u.Repository.FindByID(id)
	if err != nil {
		return nil, err
	}
	return post, nil
}
