package entity

import "github.com/tabakazu/domain-modeling-demo/domain/model/valueobject"

type Post struct {
	Model
	CustomerID valueobject.CustomerID `gorm:"embedded"`
	Name       valueobject.PostName   `gorm:"embedded"`
}

func NewPost(cusID valueobject.CustomerID, name valueobject.PostName) *Post {
	return &Post{
		CustomerID: cusID,
		Name:       name,
	}
}
