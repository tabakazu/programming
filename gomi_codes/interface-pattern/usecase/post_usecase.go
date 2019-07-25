package usecase

import (
	"github.com/tabakazu/interface-pattern/domain"
	"github.com/tabakazu/interface-pattern/infrastructure/repository"
)

type PostUsecase struct {
	PostRepository
}

type PostRepository interface {
	FindById(int) (domain.Post, error)
}

func NewPostUsecase() PostUsecase {
	return PostUsecase{
		PostRepository: repository.NewPostRepository(),
	}
}

func (u PostUsecase) GetPost(id int) (domain.Post, error) {
	post, err := u.PostRepository.FindById(id)
	if err != nil {
		return post, err
	}
	return post, nil
}
