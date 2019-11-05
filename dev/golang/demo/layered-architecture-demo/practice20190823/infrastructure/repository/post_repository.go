package repository

import (
	"github.com/tabakazu/interface-pattern-demo/domain"
	"github.com/tabakazu/interface-pattern-demo/infrastructure"
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

func (r *PostRepository) Store(post domain.Post) (domain.Post, error) {
	if err := r.DB.Conn.Create(&post).Error; err != nil {
		return post, err
	}
	return post, nil
}

func (r *PostRepository) Update(id int, postParam domain.Post) (domain.Post, error) {
	post := domain.Post{}
	if err := r.DB.Conn.First(&post, id).Error; err != nil {
		return post, err
	}
	if err := r.DB.Conn.Model(&post).Updates(postParam).Error; err != nil {
		return post, err
	}
	return post, nil
}

func (r *PostRepository) Destroy(id int) error {
	post := domain.Post{}
	if err := r.DB.Conn.First(&post, id).Error; err != nil {
		return err
	}
	if err := r.DB.Conn.Delete(&post).Error; err != nil {
		return err
	}
	return nil
}