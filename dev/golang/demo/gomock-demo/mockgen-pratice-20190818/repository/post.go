package repository

import "github.com/tabakazu/gomock-demo/domain"

type PostRepository struct{}

func (r PostRepository) FindByID(id int) (*domain.Post, error) {
	post := &domain.Post{
		ID:    id,
		Title: "Test Title",
	}
	return post, nil
}
