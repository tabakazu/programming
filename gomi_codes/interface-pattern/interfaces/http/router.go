package http

import (
	"github.com/gin-gonic/gin"
	"github.com/tabakazu/interface-pattern/interfaces/http/handler"
)

func NewRouter() *gin.Engine {
	r := gin.New()

	postHandler := handler.NewPostHandler()
	r.GET("/posts/:id", func(c *gin.Context) { postHandler.Show(c) })

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	return r
}
