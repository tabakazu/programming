package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tabakazu/layered-architecture-demo/mvc-like/model"
)

type PostsController struct{}

func NewPostsController() *PostsController {
	controller := &PostsController{}
	return controller
}

func (c *PostsController) Index(ctx *gin.Context) {
	posts := model.PostAll()
	ctx.JSON(200, posts)
}

func (c *PostsController) Show(ctx *gin.Context) {
	postID, _ := strconv.Atoi(ctx.Param("id"))
	posts := model.PostFind(postID)
	ctx.JSON(200, posts)
}
