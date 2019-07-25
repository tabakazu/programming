package repository

import (
	"github.com/tabakazu/interface-pattern/domain"
	"github.com/tabakazu/interface-pattern/infrastructure"
)

type PostRepository struct {
	DB *infrastructure.DB
}

func NewPostRepository() *PostRepository {
	return &PostRepository{
		DB: infrastructure.NewDb(),
	}
}

func (r *PostRepository) All() ([]domain.Post, error) {
	posts := []domain.Post{}
	if err := r.DB.Conn.Find(&posts).Error; err != nil {
		return posts, err
	}
	return posts, nil
}

func (r *PostRepository) FindById(id int) (domain.Post, error) {
	post := domain.Post{}
	if err := r.DB.Conn.First(&post, id).Error; err != nil {
		return post, err
	}
	return post, nil
}
