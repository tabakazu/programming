package controller

import "github.com/tabakazu/gomock-demo/usecase"

type Post struct {
	Usecase usecase.Post
}

func (c Post) Show() {
	c.Usecase.FindByID(1)
	return
}
