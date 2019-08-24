package datastore

import (
	"github.com/tabakazu/goapp-practice/domain/model"
	"github.com/tabakazu/goapp-practice/infrastructure"
)

type PostRepository struct {
	DB *infrastructure.Datastore
}

func (r *PostRepository) FindByID(id int) (*model.Post, error) {
	post := model.Post{}
	if err := r.DB.Session.First(&post, id).Error; err != nil {
		return nil, err
	}
	return &post, nil
}

func (r *PostRepository) Create(post model.Post) (*model.Post, error) {
	if err := r.DB.Session.Create(&post).Error; err != nil {
		return nil, err
	}
	return &post, nil
}
