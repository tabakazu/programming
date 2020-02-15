package database

import (
	"github.com/tabakazu/golang-api-srv/domain/model"
	"github.com/tabakazu/golang-api-srv/infrastructure"
)

type PostRepository struct {
	DbConfig infrastructure.DbConfig
}

func NewPostRepository() *PostRepository {
	return &PostRepository{
		DbConfig: infrastructure.NewDbConfig(),
	}
}

func (r *PostRepository) Find(id string) (*model.Post, error) {
	var p model.Post
	if err := r.DbConfig.Conn.First(&p, id).Error; err != nil {
		return nil, err
	}
	return &p, nil
}
