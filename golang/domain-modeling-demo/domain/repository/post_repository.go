package repository

import "github.com/tabakazu/domain-modeling-demo/domain/model"

type PostRepository interface {
	FindAll() (model.PostCollection, error)
	FindByName(name model.PostName) (model.Post, error)
}
