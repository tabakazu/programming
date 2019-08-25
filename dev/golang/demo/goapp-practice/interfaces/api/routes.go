package api

import (
	"github.com/gin-gonic/gin"
	"github.com/tabakazu/goapp-practice/interfaces/api/controller"
)

func NewRouter() *gin.Engine {
	r := gin.New()

	posts := controller.NewPostController()
	r.GET("/posts/:id", func(c *gin.Context) { posts.Show(c) })
	r.POST("/posts", func(c *gin.Context) { posts.Create(c) })

	return r
}
