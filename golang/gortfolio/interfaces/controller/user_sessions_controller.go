package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tabakazu/gortfolio/usecase"
	"github.com/tabakazu/gortfolio/usecase/requests"
)

type UserSessionsController struct {
	usecase usecase.UserLoginUseCase
}

func NewUserSessionsController(usecase usecase.UserLoginUseCase) *UserSessionsController {
	return &UserSessionsController{
		usecase: usecase,
	}
}

func (ctrl *UserSessionsController) Create(c Context) {
	req := requests.NewUserLoginRequest()

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := ctrl.usecase.Call(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}
