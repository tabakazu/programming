package repository

import (
	"github.com/tabakazu/domain-modeling-demo/domain/model/collection"
	"github.com/tabakazu/domain-modeling-demo/domain/model/entity"
	"github.com/tabakazu/domain-modeling-demo/domain/model/valueobject"
)

type PostRepository interface {
	Save(post entity.Post) (entity.Post, error)
	FindAll() (collection.PostCollection, error)
	FindByName(name valueobject.PostName) (entity.Post, error)
}
