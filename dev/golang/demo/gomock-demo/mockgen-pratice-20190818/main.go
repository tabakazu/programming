package main

import (
	"github.com/tabakazu/gomock-demo/controller"
	"github.com/tabakazu/gomock-demo/repository"
	"github.com/tabakazu/gomock-demo/usecase"
)

func main() {
	ctrl := &controller.PostController{
		Usecase: usecase.PostUsecase{
			Repository: repository.PostRepository{},
		},
	}
	ctrl.Show()
	return
}
