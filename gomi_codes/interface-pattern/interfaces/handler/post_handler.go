package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tabakazu/interface-pattern/infrastructure"
	"github.com/tabakazu/interface-pattern/infrastructure/repository"
	"github.com/tabakazu/interface-pattern/usecase"
)

type PostHandler struct {
	Usecase usecase.PostUsecase
}

func NewPostHandler(db infrastructure.DB) *PostHandler {
	return &PostHandler{
		Usecase: usecase.PostUsecase{
			PostRepository: &repository.PostRepository{
				DB: infrastructure.NewDb(),
			},
		},
	}
}

func (h *PostHandler) Show(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	post, err := h.Usecase.FindById(id)
	if err != nil {
		c.JSON(404, gin.H{"message": "Page not found"})
	} else {
		c.JSON(200, post)
	}
}
