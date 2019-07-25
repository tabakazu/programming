package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tabakazu/interface-pattern/domain"
	"github.com/tabakazu/interface-pattern/infrastructure"
	"github.com/tabakazu/interface-pattern/usecase"
)

type PostHandler struct {
	Usecase PostUsecase
}

type PostUsecase interface {
	GetPosts() ([]domain.Post, error)
	GetPost(int) (domain.Post, error)
}

func NewPostHandler(db infrastructure.DB) *PostHandler {
	return &PostHandler{
		Usecase: usecase.NewPostUsecase(),
	}
}

func (h *PostHandler) Index(c *gin.Context) {
	posts, err := h.Usecase.GetPosts()
	if err != nil {
		c.JSON(404, gin.H{"message": "Page not found"})
	} else {
		c.JSON(200, posts)
	}
}

func (h *PostHandler) Show(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	post, err := h.Usecase.GetPost(id)
	if err != nil {
		c.JSON(404, gin.H{"message": "Page not found"})
	} else {
		c.JSON(200, post)
	}
}
