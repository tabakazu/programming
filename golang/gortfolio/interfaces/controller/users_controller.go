package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tabakazu/gortfolio/usecase"
	"github.com/tabakazu/gortfolio/usecase/requests"
)

type UsersController struct {
	usecase usecase.UserCreateUseCase
}

func NewUsersController(usecase usecase.UserCreateUseCase) *UsersController {
	return &UsersController{
		usecase: usecase,
	}
}

func (ctrl *UsersController) Create(c Context) {
	req := requests.NewUserCreateRequest()

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ctrl.usecase.Call(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, nil)
}
