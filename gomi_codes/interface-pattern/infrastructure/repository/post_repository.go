package repository

import (
	"github.com/tabakazu/interface-pattern/domain"
	"github.com/tabakazu/interface-pattern/infrastructure"
)

type PostRepository struct {
	DB *infrastructure.DB
}

func (r *PostRepository) FindById(id int) (domain.Post, error) {
	post := domain.Post{}
	if err := r.DB.Conn.First(&post, id).Error; err != nil {
		return post, err
	}
	return post, nil
}
