package usecase

import (
	"github.com/tabakazu/interface-pattern/domain"
	"github.com/tabakazu/interface-pattern/infrastructure/repository"
)

type PostUsecase struct {
	PostRepository
}

type PostRepository interface {
	All() ([]domain.Post, error)
	FindById(int) (domain.Post, error)
	Store(domain.Post) (domain.Post, error)
	Update(int, domain.Post) (domain.Post, error)
}

func NewPostUsecase() PostUsecase {
	return PostUsecase{
		PostRepository: repository.NewPostRepository(),
	}
}

func (u PostUsecase) GetPosts() ([]domain.Post, error) {
	posts, err := u.PostRepository.All()
	if err != nil {
		return posts, err
	}
	return posts, nil
}

func (u PostUsecase) GetPost(id int) (domain.Post, error) {
	post, err := u.PostRepository.FindById(id)
	if err != nil {
		return post, err
	}
	return post, nil
}

func (u PostUsecase) CreatePost(post domain.Post) (domain.Post, error) {
	post, err := u.PostRepository.Store(post)
	if err != nil {
		return post, err
	}
	return post, nil
}

func (u PostUsecase) UpdatePost(id int, postParam domain.Post) (domain.Post, error) {
	post, err := u.PostRepository.Update(id, postParam)
	if err != nil {
		return post, err
	}
	return post, nil
}
