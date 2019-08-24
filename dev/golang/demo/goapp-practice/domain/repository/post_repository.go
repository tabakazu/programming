package repository

import "github.com/tabakazu/goapp-practice/domain/model"

type PostRepository interface {
	FindByID(id int) (*model.Post, error)
	Create(post model.Post) (*model.Post, error)
}
