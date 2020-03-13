package model

type Post struct {
	Name PostName
}

func NewPost(name PostName) *Post {
	return &Post{
		Name: name,
	}
}
