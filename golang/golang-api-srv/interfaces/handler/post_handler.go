package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tabakazu/golang-api-srv/usecase"
)

type PostHandler struct {
	Usecase *usecase.PostUsecase
}

func NewPostHandler() *PostHandler {
	return &PostHandler{
		Usecase: usecase.NewPostUsecase(),
	}
}

func (h *PostHandler) Show(c *gin.Context) {
	p, err := h.Usecase.GetPost(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, p)
}
