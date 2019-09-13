package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/tabakazu/layered-architecture-demo/mvc-like/controller"
)

func NewRoutes() *gin.Engine {
	r := gin.New()

	posts := controller.NewPostsController()
	r.GET("/posts", func(c *gin.Context) { posts.Index(c) })
	r.GET("/posts/:id", func(c *gin.Context) { posts.Show(c) })

	return r
}
