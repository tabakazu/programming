package controller

import (
	"fmt"

	"github.com/tabakazu/gomock-demo/domain"
)

type PostUsecase interface {
	GetPostByID(id int) (*domain.Post, error)
}

type PostController struct {
	Usecase PostUsecase
}

func (c PostController) Show() {
	post, err := c.Usecase.GetPostByID(1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(post)
	return
}
