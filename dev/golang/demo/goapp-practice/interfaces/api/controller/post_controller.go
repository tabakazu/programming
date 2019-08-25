package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tabakazu/goapp-practice/application/usecase"
	"github.com/tabakazu/goapp-practice/domain/model"
)

type PostController struct {
	Usecase *usecase.PostUsecase
}

func NewPostController() *PostController {
	return &PostController{
		Usecase: usecase.NewPostUsecase(),
	}
}

func (ctrl *PostController) Show(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(500, gin.H{"message": "500 Internal Server Error"})
		return
	}

	post, err := ctrl.Usecase.GetPostByID(id)
	if err != nil {
		c.JSON(404, gin.H{"message": "404 Not Found"})
		return
	}

	c.JSON(200, post)
	return
}

func (ctrl *PostController) Create(c *gin.Context) {
	param := model.Post{}
	if err := c.BindJSON(&param); err != nil {
		c.JSON(422, gin.H{"message": "422 Unprocessable Entity"})
		return
	}

	post, err := ctrl.Usecase.CreatePost(param)
	if err != nil {
		c.JSON(422, gin.H{"message": "422 Unprocessable Entity"})
		return
	}

	c.JSON(201, post)
	return
}
