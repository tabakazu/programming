package domain

import "github.com/jinzhu/gorm"

type Post struct {
	gorm.Model
	Title string `gorm:"not null" form:"title" json:"title"`
	Body  string `gorm:"not null" form:"body" json:"body"`
}
