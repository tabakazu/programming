package datastore

import (
	"github.com/tabakazu/domain-modeling-demo/domain/model/collection"
	"github.com/tabakazu/domain-modeling-demo/domain/model/entity"
	"github.com/tabakazu/domain-modeling-demo/domain/model/valueobject"
	"github.com/tabakazu/domain-modeling-demo/domain/repository"
)

type PostRepository struct {
	*Config
}

func NewPostRepository() repository.PostRepository {
	config := NewConfig()
	return &PostRepository{Config: config}
}

func (r *PostRepository) Save(post entity.Post) (entity.Post, error) {
	if err := r.Conn.Create(&post).Error; err != nil {
		return post, err
	}
	return post, nil
}

func (r *PostRepository) FindAll() (collection.PostCollection, error) {
	var postCollection collection.PostCollection
	if err := r.Conn.Find(&postCollection.List).Error; err != nil {
		return postCollection, err
	}
	return postCollection, nil
}

func (r *PostRepository) FindByName(name valueobject.PostName) (entity.Post, error) {
	var post entity.Post
	if err := r.Conn.Where("name = ?", name.Value).First(&post).Error; err != nil {
		return post, err
	}
	return post, nil
}
