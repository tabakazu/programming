package model

import "time"

type Post struct {
	ID        uint   `gorm:"primary_key" gorm:"not null" form:"id" json:"id"`
	Title     string `gorm:"not null" form:"title" json:"title"`
	Body      string `gorm:"not null" form:"body" json:"body"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
