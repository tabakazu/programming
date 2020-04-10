package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tabakazu/gortfolio/usecase"
	"github.com/tabakazu/gortfolio/usecase/requests"
)

type UsersController struct {
	usecase usecase.UserRegisterUseCase
}

func NewUsersController(usecase usecase.UserRegisterUseCase) *UsersController {
	return &UsersController{
		usecase: usecase,
	}
}

func (ctrl *UsersController) Create(c Context) {
	req := requests.NewUserRegisterRequest()

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ctrl.usecase.Call(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusCreated)
}
