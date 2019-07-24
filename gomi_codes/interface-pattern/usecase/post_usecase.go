package usecase

import "github.com/tabakazu/interface-pattern/domain"

type PostUsecase struct {
	PostRepository
}

type PostRepository interface {
	FindById(int) (domain.Post, error)
}

func (u PostUsecase) FindById(id int) (domain.Post, error) {
	post, err := u.PostRepository.FindById(id)
	if err != nil {
		return post, err
	}
	return post, nil
}
