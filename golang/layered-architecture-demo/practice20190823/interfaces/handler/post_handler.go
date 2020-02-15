package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tabakazu/interface-pattern-demo/domain"
	"github.com/tabakazu/interface-pattern-demo/infrastructure"
	"github.com/tabakazu/interface-pattern-demo/usecase"
)

type PostHandler struct {
	Usecase PostUsecase
}

type PostUsecase interface {
	GetPosts() ([]domain.Post, error)
	GetPost(int) (domain.Post, error)
	CreatePost(domain.Post) (domain.Post, error)
	UpdatePost(int, domain.Post) (domain.Post, error)
	DestroyPost(int) error
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
	return
}

func (h *PostHandler) Show(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	post, err := h.Usecase.GetPost(id)
	if err != nil {
		c.JSON(404, gin.H{"message": "Page not found"})
	} else {
		c.JSON(200, post)
	}
	return
}

func (h *PostHandler) Create(c *gin.Context) {
	post := domain.Post{}
	c.BindJSON(&post)
	post, err := h.Usecase.CreatePost(post)
	if err != nil {
		c.JSON(404, gin.H{"message": "Page not found"})
	} else {
		c.JSON(201, post)
	}
	return
}

func (h *PostHandler) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	postParam := domain.Post{}
	c.BindJSON(&postParam)
	post, err := h.Usecase.UpdatePost(id, postParam)
	if err != nil {
		c.JSON(404, gin.H{"message": "Page not found"})
	} else {
		c.JSON(200, post)
	}
	return
}

func (h *PostHandler) Destroy(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.Usecase.DestroyPost(id); err != nil {
		c.JSON(404, gin.H{"message": "Page not found"})
		return
	}
	c.JSON(200, gin.H{"message": "Successful"})
	return
}
