package usecase

import (
	"github.com/tabakazu/golang-api-srv/domain/model"
	"github.com/tabakazu/golang-api-srv/infrastructure/database"
)

type PostUsecase struct {
	Repo *database.PostRepository
}

func NewPostUsecase() *PostUsecase {
	return &PostUsecase{
		Repo: database.NewPostRepository(),
	}
}

func (u *PostUsecase) GetPost(id string) (*model.Post, error) {
	p, err := u.Repo.Find(id)
	if err != nil {
		return nil, err
	}
	return p, nil
}
