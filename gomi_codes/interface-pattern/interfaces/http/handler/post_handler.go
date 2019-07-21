package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type PostHandler struct{}

func NewPostHandler() *PostHandler {
	return &PostHandler{}
}

func (handler *PostHandler) Show(c *gin.Context) {
	fmt.Println((c.Param("id")))
}
