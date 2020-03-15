package entity

import "github.com/tabakazu/domain-modeling-demo/domain/model/valueobject"

type Post struct {
	Name valueobject.PostName `gorm:"embedded"`
}

func NewPost(name valueobject.PostName) *Post {
	return &Post{
		Name: name,
	}
}
