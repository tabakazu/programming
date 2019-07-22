package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/tabakazu/interface-pattern/domain"
	"github.com/tabakazu/interface-pattern/infrastructure"
)

type PostHandler struct {
	DB infrastructure.DB
}

func NewPostHandler(db infrastructure.DB) *PostHandler {
	return &PostHandler{
		DB: db,
	}
}

func (handler *PostHandler) Show(c *gin.Context) {
	post := domain.Post{}
	if err := handler.DB.Conn.First(&post, c.Param("id")).Error; err != nil {
		c.JSON(200, gin.H{
			"message": "NotFoundRecord",
		})
	} else {
		c.JSON(200, post)
	}
}
