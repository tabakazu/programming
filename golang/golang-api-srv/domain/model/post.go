package model

type Post struct {
	Model
	Title string `json:"title" binding:"required"`
	Body  string `json:"body" binding:"required"`
}
