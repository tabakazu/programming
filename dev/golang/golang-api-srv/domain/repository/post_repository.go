package repository

import "github.com/tabakazu/golang-api-srv/domain/model"

type PostRepository interface {
	Find(id string) (*model.Post, error)
}
