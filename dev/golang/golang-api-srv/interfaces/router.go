package interfaces

import (
	"github.com/gin-gonic/gin"
	"github.com/tabakazu/golang-api-srv/interfaces/handler"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	postHandler := handler.NewPostHandler()
	r.GET("/posts/:id", postHandler.Show)
	return r
}
