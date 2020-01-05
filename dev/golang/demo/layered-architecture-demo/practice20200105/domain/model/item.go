package model

type Item struct {
	Model
	Title string `gorm:"column:title" json:"title"`
}
